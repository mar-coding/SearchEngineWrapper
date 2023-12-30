#!/bin/bash

# Load environment variables from the '.env' file
if [ -f .env ]; then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

# Create the 'config' directory if it does not already exist
if [ ! -d "$PARENT_DIR/conf" ]; then
  mkdir "$PARENT_DIR/conf"
  mkdir "$PARENT_DIR/conf/esdata01"
  mkdir "$PARENT_DIR/conf/kibanadata"
fi

# Change the ownership of the 'config' directory and all its subdirectories and files
sudo chown -R $USER:$USER "$PARENT_DIR/conf"
sudo chown -R $USER:$USER "$PARENT_DIR/conf/esdata01"
sudo chmod 755 -R "$PARENT_DIR/conf" 