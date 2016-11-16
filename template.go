// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"io/ioutil"
	"path"
	"time"

	"gopkg.in/yaml.v2"
)

type template struct {
	Peers           []string          `yaml:"peers"`
	Peer            string            `yaml:"peer"`
	PeerList        string            `yaml:"peerlist"`
	Caller          string            `yaml:"caller"`
	Service         string            `yaml:"service"`
	Thrift          string            `yaml:"thrift"`
	Procedure       string            `yaml:"procedure"`
	ShardKey        string            `yaml:"shardkey"`
	RoutingKey      string            `yaml:"routingkey"`
	RoutingDelegate string            `yaml:"routingdelegate"`
	Headers         map[string]string `yaml:"headers"`
	Baggage         map[string]string `yaml:"baggage"`
	Jaeger          bool              `yaml:"jaeger"`
	Request         interface{}       `yaml:"request"`
	Timeout         time.Duration     `yaml:"timeout"`
}

func readYamlRequest(opts *Options) error {
	t := template{}

	bytes, err := ioutil.ReadFile(opts.ROpts.YamlTemplate)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bytes, &t)
	if err != nil {
		return err
	}

	body, err := yaml.Marshal(t.Request)
	if err != nil {
		return err
	}

	if t.Peer != "" {
		opts.TOpts.HostPorts = []string{t.Peer}
	} else if t.Peers != nil {
		opts.TOpts.HostPorts = t.Peers
	} else if t.PeerList != "" {
		opts.TOpts.HostPortFile = path.Join(path.Dir(opts.ROpts.YamlTemplate), t.PeerList)
	}

	// Header arguments have precedence over template headers.
	if t.Headers != nil {
		if opts.ROpts.Headers == nil {
			opts.ROpts.Headers = t.Headers
		} else {
			for k, v := range t.Headers {
				if _, exists := opts.ROpts.Headers[k]; !exists {
					opts.ROpts.Headers[k] = v
				}
			}
		}
	}

	if t.Jaeger {
		opts.TOpts.Jaeger = true
	}

	// Baggage arguments have precedence over template headers.
	if t.Baggage != nil {
		if opts.ROpts.Baggage == nil {
			opts.ROpts.Baggage = t.Baggage
		} else {
			for k, v := range t.Baggage {
				if _, exists := opts.ROpts.Baggage[k]; !exists {
					opts.ROpts.Baggage[k] = v
				}
			}
		}
	}

	opts.ROpts.ThriftFile = t.Thrift
	opts.TOpts.CallerName = t.Caller
	opts.TOpts.ServiceName = t.Service
	opts.ROpts.MethodName = t.Procedure
	opts.TOpts.ShardKey = t.ShardKey
	opts.TOpts.RoutingKey = t.RoutingKey
	opts.TOpts.RoutingDelegate = t.RoutingDelegate
	opts.ROpts.RequestJSON = string(body)
	opts.ROpts.Timeout = timeMillisFlag(t.Timeout)

	return nil
}
