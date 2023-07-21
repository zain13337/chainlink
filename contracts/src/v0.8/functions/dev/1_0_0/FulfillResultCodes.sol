// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

enum FulfillResult {
  USER_SUCCESS, // 0
  USER_ERROR, // 1
  INVALID_REQUEST_ID, // 2
  INSUFFICIENT_GAS, // 3
  INSUFFICIENT_SUBSCRIPTION_BALANCE, // 4
  COST_EXCEEDS_COMMITMENT, // 5
  INTERNAL_ERROR // 6
}