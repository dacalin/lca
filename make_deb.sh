#!/bin/bash
VERSION="1.0.1"  # Set version

dpkg-deb --build --root-owner-group linux/amd64 linux/deb/lca-${VERSION}-amd64.deb
dpkg-deb --build --root-owner-group linux/arm64 linux/deb/lca-${VERSION}-arm64.deb
cp linux/deb/lca-${VERSION}-amd64.deb linux/deb/lca-latest-amd64.deb
cp linux/deb/lca-${VERSION}-arm64.deb linux/deb/lca-latest-arm64.deb