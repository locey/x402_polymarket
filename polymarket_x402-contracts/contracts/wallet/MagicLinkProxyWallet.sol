// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title MagicLinkProxyWallet
 * @dev 代理钱包实现合约（逻辑合约）
 */
contract MagicLinkProxyWallet {
    address public owner;
    address public factory;
    uint256 public nonce;

    event TranscationExecute(address indexed target, uint256 value, bytes data, bytes result);
    modifier onlyOwner(){
        require(msg.sender == owner, "Only owner can call");
        _;
    }

    function initialize(address _owner, address _factory) external {
        require(owner == address(0), "Already initialized");
        require(_owner != address(0), "Invalid owner");
        require(_factory != address(0), "Invalid factory");
        owner = _owner;
        factory = _factory;
    }

    function execute(address target, uint256 value, bytes calldata data) external onlyOwner returns (bytes memory){
        require(target != address(0), "Invalid target");
        nonce++;
        (bool success, bytes memory result) = target.call{value: value}(data);
        require(success, "Transaction failled");
        emit TranscationExecute(target, value, data, result);
        return result;
    }
    // 接收 ETH
    receive() external payable{}
}