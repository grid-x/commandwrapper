#!/bin/bash

kill -SIGTERM $(ps -a | grep "commandwrapper" | grep -v grep | awk '{print $1}')