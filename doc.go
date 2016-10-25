// This file is part of graze/golang-service
//
// Copyright (c) 2016 Nature Delivered Ltd. <https://www.graze.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.
//
// license: https://github.com/graze/golang-service/blob/master/LICENSE
// link:    https://github.com/graze/golang-service

/*
golang-service is a set of packages to help with creating services using golang for logging and testing

golang-service contains the following packages:

The logging package provides some logging helpers for statsd, logfmt and syslog

The handlers package provides a set of handlers that handle http.Request log the results

The nettest package provides a set of helpers for use when testing networks
*/
package golangservice

import (
	_ "github.com/graze/golang-service/handlers"
	_ "github.com/graze/golang-service/logging"
	_ "github.com/graze/golang-service/nettest"
)
