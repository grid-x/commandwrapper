#!/bin/bash

kill -SIGTERM $(ps | grep "commandwrapper" | grep -v grep | awk '{print $1}')