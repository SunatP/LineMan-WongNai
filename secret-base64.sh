#!/bin/bash
secret="aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
echo $secret |base64 -d | rev
