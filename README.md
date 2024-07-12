# Tinyval Go

Hash table service for storing and retrieving values (blobs) using their SHA-256 hashes as keys.

This service could be deployed as many single-node instances each using own local storage (e.g. a filesystem).
Responsibilities involving high availability, partitioning, per-value authorization, and anything metadata related are delegated to other services.
The idea is to let this service be a general-purpose component used to build distributed hash table (DHT) based systems.

Note that this is primarily an experiment in using Go to develop HTTP services.
