// Copyright (c) 2016 Uber Technologies, Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
namespace java com.uber.cherami

include "shared.thrift"

enum NotificationType {
  HOST = 1,
  CLIENT = 2,
  ALL = 3
}

struct ReconfigureClientInfo {
  1: optional i32 numberOfConnections
}

struct DestinationUpdatedNotification {
  1: optional string destinationUUID
  2: optional NotificationType type
  3: optional ReconfigureClientInfo clientInfo
  4: optional string extentUUID
  5: optional list<string> storeIds
}

struct DestinationsUpdatedRequest {
  1: optional string updateUUID
  2: optional list<DestinationUpdatedNotification> updates
}

struct UnloadDestinationsRequest {
  1: optional list<string> destUUIDs
}

struct Destinations {
  1: optional string destUUID
  2: optional string destPath
}

struct ListLoadedDestinationsResult {
  1: optional list<Destinations> dests
}

struct InputDestExtent {
  1: optional string extentUUID
  2: optional i64 maxSeqNo
  3: optional i64 maxSizeBytes
  4: optional i64 currSeqNo
  5: optional i64 currSizeBytes
  6: optional list<string> replicas
}

struct DestinationState {
  1: optional string destUUID
  2: optional i64 msgsChSize
  3: optional i64 numConnections
  4: optional i64 numMsgsIn
  5: optional i64 numSentAcks
  6: optional i64 numSentNacks
  7: optional i64 numThrottled
  8: optional i64 numFailed
  9: optional list<InputDestExtent> destExtents
}

struct ReadDestinationStateResult {
  1: optional string inputHostUUID
  2: optional list<DestinationState> destState
}

struct ReadDestinationStateRequest {
  1: optional list<string> destUUIDs
}

service InputHostAdmin {
  void destinationsUpdated(1: DestinationsUpdatedRequest request)
  void unloadDestinations(1: UnloadDestinationsRequest request)
  ListLoadedDestinationsResult listLoadedDestinations()
  ReadDestinationStateResult readDestState(1: ReadDestinationStateRequest request)
  /*
   * start sealing in inputhost. As soon as an inputhost receives this call, it will
   * do the following:
   * (1) Start draining the specified extent gracefully
   * (2) Notify clients to drain their connections as well, as necessary (i.e, if this is 
   *     the only extent hosted on this inputhost)
   * (3) Once drained notify the Extent Controller using the ExtentsUnreachable() API to
   *     actually seal the extent on the storehosts.
   */
  void sealExtent(1: shared.SealExtentRequest sealRequest) // start seal in inputhost. This will drain this extent and
    throws (
      1: shared.ExtentSealedError sealedError,
      2: shared.ExtentFailedToSealError failedError)
}

struct ConsumerGroupUpdatedNotification {
  1: optional string consumerGroupUUID
  2: optional NotificationType type
  3: optional ReconfigureClientInfo clientInfo
}

struct ConsumerGroupsUpdatedRequest {
  1: optional string updateUUID
  2: optional list<ConsumerGroupUpdatedNotification> updates
}

struct UnloadConsumerGroupsRequest {
  1: optional list<string> cgUUIDs
}

struct AckMgrState {
  1: optional i16 ackMgrID
  2: optional bool isSealed
  3: optional i64 readLevelSeq
  4: optional i64 ackLevelSeq
  5: optional i64 readLevelOffset
  6: optional i64 ackLevelOffset
  7: optional i64 lastAckLevelUpdateTime
  8: optional i64 lastAckedSeq
  9: optional i64 numAckedMsgs
  10: optional i64 numUnackedMsgs
}

struct OutputCgExtent {
  1: optional string extentUUID
  2: optional string connectedStoreUUID
  3: optional i32 numCreditsSentToStore
  4: optional i32 numMsgsReadFromStore
  5: optional i64 startSequence
  6: optional AckMgrState ackMgrState
}
 
struct ConsumerGroupState {
  1: optional string cgUUID
  2: optional i32 numOutstandingMsgs
  3: optional i64 msgChSize
  4: optional i64 msgCacheChSize
  5: optional i64 numConnections
  6: optional list<OutputCgExtent> cgExtents
}

struct ReadConsumerGroupStateResult {
  1: optional i16 sessionID
  2: optional list<ConsumerGroupState> cgState
}

struct ReadConsumerGroupStateRequest {
  1: optional list<string> cgUUIDs
}

struct ConsumerGroups {
  1: optional string cgUUID
  2: optional string cgName
  3: optional string destPath
}

struct ListConsumerGroupsResult {
  1: optional list<ConsumerGroups> cgs
}
 
service OutputHostAdmin {
  void consumerGroupsUpdated(1: ConsumerGroupsUpdatedRequest request)
  void unloadConsumerGroups(1: UnloadConsumerGroupsRequest request)
  ReadConsumerGroupStateResult readCgState(1: ReadConsumerGroupStateRequest request)
  ListConsumerGroupsResult listLoadedConsumerGroups()
}

struct ExtentUnreachableNotification {
  1: optional string destinationUUID
  2: optional string extentUUID
  3: optional i64 (js.type = "Long") sealSequenceNumber
}

struct ExtentsUnreachableRequest {
  1: optional string updateUUID
  2: optional list<ExtentUnreachableNotification> updates
}

struct ExtentDrainNotification {
  1: optional string destinationUUID
  2: optional string extentUUID
}

struct ExtentDrainRequest {
  1: optional string updateUUID
  2: optional list<ExtentDrainNotification> updates
}

service ControllerHostAdmin {
  /*
   * This is one of the triggers to seal an extent on the store host
   * by the extent controller.
   * InputHost can call this API to immediately seal an extent.
   * When an extent controller recieves this request, it should call the
   * SealExtent API on the store host to seal at the specified sequence
   * number, if available. If the sequence number is not given, the
   * extent controller should find the least possible sequence number from
   * all the stores and then seal the extent.
   */
  void extentsUnreachable(1: ExtentsUnreachableRequest request)
  /*
   * This is one of the triggers to start a graceful drain for the extent.
   * This is used by the inputhost and/or the storehost to initiate a
   * graceful drain for this extent (or a bunch of extents).
   * For example: If an inputhost notices an extent is beyond it's hard limit and
   * has not started draining yet, it will use this API to signal the controller which will
   * in turn start the drain as usual as soon as it sees this message.
   * This can also be used during upgrades/deployments to gracefully drain extents.
   */
  void initiateDrain(1: ExtentDrainRequest request)
}
