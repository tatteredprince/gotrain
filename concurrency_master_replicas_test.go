package main

import (
	"context"
	"errors"
	"time"
)

type Client interface {
	Get(ctx context.Context, key string) (result string, err error)
}

// requestMasterReplicas requests master and two replicas for key's value.
// If master responded within its timeout then its response is returned.
// If request to master timeouted or is erroneous then requests are sent to replicas.
// If one of the requests to the replicas responded with error then that error is returned.
// If master responded after requests are sent to replicas then that response is returned.
func requestMasterReplicas(
	ctx context.Context,
	master Client,
	replicas [2]Client,
	key string,
	masterTimeout time.Duration,
	overallTimeout time.Duration,
) (string, error) {
	masterTimer := time.After(masterTimeout)
	overallTimer := time.After(overallTimeout)
	respCh := make(chan string)
	errCh := make(chan error)

	send := func(client Client) {
		if response, err := client.Get(ctx, key); err != nil {
			errCh <- err
		} else {
			respCh <- response
		}
	}
	sendToReplicas := func(clients [2]Client) (string, error) {
		go send(clients[0])
		go send(clients[1])
		select {
		case response := <-respCh:
			return response, nil
		case err := <-errCh:
			return "", err
		case <-overallTimer:
			return "", errors.New("requests took too long")
		}
	}

	go send(master)

	select {
	case response := <-respCh:
		return response, nil
	case <-errCh:
		if resp, err := sendToReplicas(replicas); err != nil {
			return "", err
		} else {
			return resp, nil
		}
	case <-masterTimer:
		if resp, err := sendToReplicas(replicas); err != nil {
			return "", err
		} else {
			return resp, nil
		}
	case <-overallTimer:
		return "", errors.New("requests took long")
	}
}
