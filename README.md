# Plugin Development Tutorial

This repository provides a tutorial for developing new plugins for Blueprint. The repository covers three major use-cases for plugins:

+ Instrumenting services in an application's workflow
+ Adding new methods/APIs to services in an application's workflow
+ Modifying the function signatures of methods for services in an application's workflow

The repository provides a plugin implementation for each of the aforementioned types along with a simple two-service application to demonstrate how to use the plugins in an application. The repository is structured as follows:

+ [plugins/tutorial](plugins/tutorial): Package that contains the implementation of the `tutorial` plugins.
+ [examples/helloworld](examples/helloworld): Module that provides the two-service application and demonstrates the use of the `tutorial` plugins.
