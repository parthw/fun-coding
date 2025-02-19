#!/bin/sh
sudo apt-get update
sudo apt-get install -y git gcc make curl

pip install --upgrade pip
pip install numpy pandas scipy matplotlib seaborn scikit-learn torch requests plotly jupyterlab_git certifi setuptools wheel