#!/bin/bash

find . -name "*.*" -type f -exec cat {} + | xclip -selection clipboard
