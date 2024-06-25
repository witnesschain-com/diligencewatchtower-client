# op-watchtower

> The First Line of Defence for Optimistic Rollups

## Introduction

This is the repository for the watchtower service for Optimistic Rollups (ORs) provided by WitnessChain. This document entails the goals, our relevant learnings about the OP-Stack, our understanding about the problem and our approach to various challenges that may arise. This is likely to update overtime.

## What to watch

Our first version of the `op-watchtower` will be watching the L1(ethereum)'s L2 output oracle contract of optimism chain for the `OutputProposed` event logs, the output root from these are matched against the state root from `op-geth` of our node setup.

## How to watch

## How to raise alarms

Currently this is TBD, some possible options are:

- on L1 chain in an existing important contract (not sure which one, possibly L2OO)
- on L2 chain in an existing important contract
- on L1 chain in our own custom contract
- on L2 chain in our own custom contract
- on our own off-chain dashboard

## How to rotate watchtowers

Currently this is TBD
