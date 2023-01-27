package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	WithTimeout()

}

func WithTimeout() {

	//Attaching to etcdkeeper, golang-study-etcd-1, golang-study-jaeger-all-in-one-1, golang-study-mysql8-1, golang-study-otel-collector-1, golang-study-redis-1, golang-study-zipkin-all-in-one-1, prometheus
	//Error response from daemon: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: error during container init: error mounting "/run/desktop/mnt/host/e/Code/golange/golang-study/otel-collector-config.yaml" to rootfs at "/etc/otel-collector-config.yaml": mou
	//nt /run/desktop/mnt/host/e/Code/golange/golang-study/otel-collector-config.yaml:/etc/otel-collector-config.yaml (via /proc/self/fd/14), flags: 0x5000: not a directory: unknown: Are you trying to mount a directory onto a file (or vice-versa)? Check if the specified host path exists and is the expected type
	context, cancel := context.WithTimeout(context.TODO(), time.Second*5)

	defer cancel()
	for {
		select {
		case <-context.Done():
			fmt.Println("完成！")
			return
		default:
			fmt.Println("default!")
		}
	}

}
