# gounite[![Build Status](https://travis-ci.com/muguangyi/gounite.svg?branch=developer)](https://travis-ci.com/muguangyi/gounite)

**Gounite** 是`golang`实现的一套轻量级服务器开发框架，以**单元依赖**为规则建立容器互联，允许用户灵活定制自己的服务器架构，并能快速建立易于扩展的服务器开发解决方案。

## 框架

    +----------------------------+            +=======+  register  +------------+
    | union                      |  register  |       |<<<<<<>>>>>>| union      |
    |                            |<<<<<<>>>>>>|       |   query    +------------+
    |                            |   query    |  hub  |
    |                            |            |       |  register  +------------+
    | +------------------------+ |            |       |<<<<<<>>>>>>| union      |
    | | unit 1                 | |            +=======+   query    |            |
    | | unit 2 (depend unit N) | |                                 | +--------+ |
    | +------------------------+ |<------------------------------->| | unit N | |
    |                            |        directly connected       | +--------+ |
    +----------------------------+                                 +------------+

## 技术点

* 一个单元容器（union）是一个独立的服务器节点，可容纳多个功能单元(unit)
* 每一个功能单元（unit）运行在一个独立协程中
* 单元与单元间通过管道RPC通信（暂时只提供`同步`方式）
* 不同容器内的单元也可以通过同样的方式通信（建立在**单元依赖**的容器互联）
* **容器基站**：提供服务注册和查询功能，每一个容器都要向至少一个基站注册。

### Unit

所用的功能都是一个独立的`Unit`，并且`Unit`之间的调用无需关心是在同一个容器中，还是不同的容器，是在同一个物理机还是在不同的物理机。

### Union

逻辑单元容器，是一个独立的服务器节点（可以是一个独立的服务器进程，也允许多个容器在一个服务器进程）。

### Hub

单元容器基站，为单元容器分配端口，管理各种单元的注册和发现功能。