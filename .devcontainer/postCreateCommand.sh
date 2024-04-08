#!/bin/bash

# .zshrcのパス
ZSHRC_PATH=~/.zshrc

# 環境変数を追加する関数
add_env_variable() {
  if grep -qxF "export $1" $ZSHRC_PATH; then
    echo "$1 already exists in $ZSHRC_PATH"
  else
    echo "export $1" >> $ZSHRC_PATH
    echo "Added $1 to $ZSHRC_PATH"
  fi
}

# 環境変数を追加する
add_env_variable "DB_USER=docker"
add_env_variable "DB_PASSWORD=password"
add_env_variable "DB_HOST=mysql"
add_env_variable "DB_PORT=3306"
add_env_variable "DB_NAME=template"
