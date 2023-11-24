# Description

== A Distributed Auction System ==

::Introduction::

You must implement a **distributed auction system** using replication: a distributed component which handles auctions, and provides operations for bidding and querying the state of an auction. The component must faithfully implement the semantics of the system described below, and must at least be resilient to one (1) crash failure.

::MA Learning Goal::

The goal of this mandatory activity is that you learn (by doing) how to use replication to design a service that is resilent to crashes. In particular, it is important that you can recognise what the key issues that may arise are and understand how to deal with them.

::API::

Your system must be implemented as some number of nodes,  running on distinct processes (no threads). Clients direct API requests to *any* node they happen to know (it is up to you to decide how many nodes can be known). Nodes must respond to the following API

Method:  bid

Inputs:  amount (an int)

Outputs: ack

Comment: given a bid, returns an outcome among {fail, success or exception}

Method:  result

Inputs:  void

Outputs: outcome

Comment:  if the auction is over, it returns the result, else highest bid.

::Semantics::

Your component must have the following behaviour, for any reasonable sequentialisation/interleaving of requests to it:

- The first call to "bid" registers the bidder.
- Bidders can bid several times, but a bid must be higher than the previous one(s).
- after a predefined timeframe, the highest bidder ends up as the winner of the auction, e.g, after 100 time units from the start of the system.
- bidders can query the system in order to know the state of the auction.

:: Faults ::

- Assume a network that has reliable, ordered message transport, where transmissions to non-failed nodes complete within a known time-limit.
- Your component must be resilient to the failure-stop failure of one (1) node.

# Requirement Engineering

Implement a distributed auction system:

### Functional Requirements

### R1: Auction Bidding

- **R1.1:** The system must allow registered bidders to place bids.
- **R1.2:** A bid must be an integer value.
- **R1.3:** Each bid must be higher than any previous bid made by the same bidder.
- **R1.4:** The first call to "bid" registers the bidder.

### R2: Auction Outcome

- **R2.1:** The system must provide the current highest bid when requested.
- **R2.2:** Once the auction is over, the system must declare the highest bidder as the winner.
- **R2.3:** The auction ends after a predefined time frame (for sake of simplity, auction ends after a total of 5 bids were placed).

### R3: Distributed Processing

- **R3.1:** The system must consist of multiple nodes running on separate processes.
- **R3.2:** Clients can direct API requests to any node they are aware of.

### Non-Functional Requirements

### R4: Fault Tolerance

- **R4.1:** The system must tolerate the failure-stop failure of one Replica Manager (RM) node, in particular the Primary RM node.
- **R4.2:** The system should continue functioning with the remaining nodes after a failure.

### R5: Replication

- **R6.1:** Implement a leader-based passive replication.
- **R6.2:** There should be a primary Replica Manager (RM) and three backup RMs.
- **R6.3:** The system must handle the failure of the Primary RM and elect a new Primary RM.

### R7: Client-Server Interaction

- **R7.1:** The system must provide a front-end API for client interactions.
- **R7.2:** Clients should be able to determine the current Primary RM.

### Implementation Requirements

### R8: Development Environment

- **R8.1:** Implement the system in GoLang.

### R9: Documentation and Logging

- **R9.1:** Provide a comprehensive README file with instructions for running the implementation.
- **R9.2:** Submit logs documenting a correct system run, including handling failures.

### Phase Requirements

### R10: Auction Phases

- **R10.1:** The system must distinctly separate the bidding phase and the outcome query phase.
- **R10.2:** There must be a clear timestamp or mechanism to differentiate between the two phases.