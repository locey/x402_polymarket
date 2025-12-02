// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/proxy/Clones.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import "./MagicLinkProxyWallet.sol";

contract MagicLinkProxyFactory is Ownable, EIP712 {
    using Clones for address;
    using ECDSA for bytes32;

    address public immutable implementation;
    // MagicLink 验证者地址（签名验证）
    address public magicLinkVerifier;
    // 用户地址 => 代理钱包地址
    mapping(address => address) public userToProxy;
    // 代理钱包地址 => 是否已初始化
    mapping(address => bool) public proxyInitialized;

    event ProxyCreated(address indexed proxy, address indexed user, uint256 timestamp);
    event VerifierUpdated(address indexed oldVerifier, address indexed newVerifier);

    // EIP712 域名配置
    bytes32 public constant CREATE_PROXY_TYPEHASH = keccak256(
        "CreateProxy(address user,uint256 nonce)"
    );

    constructor(
        address _implementation, 
        address _magicLinkVerifier
    ) EIP712("MagicLinkProxyFactory", "1") Ownable(msg.sender){
        require(_implementation != address(0), "Invalid implementation");
        require(_magicLinkVerifier != address(0), "Invalid verifier");
        implementation = _implementation;
        magicLinkVerifier = _magicLinkVerifier;
    }

    function createProxyWithMagicLink(address user, uint256 nonce, bytes calldata signature) 
        external returns (address proxy){
            require(userToProxy[user] != address(0), "proxy already exists");
            require(user != address(0), "Invalid user");
            //验证签名
            bytes32 digest = _hashTypedDataV4(keccak256(
                abi.encode(CREATE_PROXY_TYPEHASH, user, nonce)
            ));

            address signer = digest.recover(signature);
            require(signer == magicLinkVerifier, "Invalid signature");

            //创建EIP-1167最小代理
            proxy = implementation.clone();

            //初始化代理钱包
            MagicLinkProxyWallet(proxy).initialize(user, address(this));

            //更新状态
            userToProxy[user] = proxy;
            proxyInitialized[proxy] = true;

            emit ProxyCreated(proxy, user, block.timestamp);
    }

    function updateMagicLinkVerifier(address _newVerifier) external onlyOwner {
        require(_newVerifier != address(0), "Invalid new verifier");
        emit VerifierUpdated(magicLinkVerifier, _newVerifier);
        magicLinkVerifier = _newVerifier;
    }
    
    function getProxyByUser(address user) external view returns (address) {
        return userToProxy[user];
    }
}