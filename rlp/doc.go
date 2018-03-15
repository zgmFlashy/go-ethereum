// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

/*
Package rlp implements the RLP serialization format.

The purpose of RLP (Recursive Linear Prefix) is to encode arbitrarily
nested arrays of binary data, and RLP is the main encoding method used
to serialize objects in Ethereum. The only purpose of RLP is to encode
structure; encoding specific atomic data types (eg. strings, ints,
floats) is left up to higher-order protocols; in Ethereum integers
must be represented in big endian binary form with no leading zeroes
(thus making the integer value zero equivalent to the empty byte
array).

RLP values are distinguished by a type tag. The type tag precedes the
value in the input stream and defines the size and kind of the bytes
that follow.
*/

/*
包rlp实现RLP序列化格式。

RLP（递归线性前缀）的目的是任意编码
嵌套的二进制数据数组，而RLP是使用的主要编码方法
在Ethereum中序列化对象。 RLP的唯一目的是编码
结构体; 编码特定的原子数据类型（例如，字符串，整数，
浮动）是由高阶协议决定的; 在以太坊整数
必须以大端二进制形式表示，且不含前导零
（从而使整数值零等于空字节
数组）。

RLP值由类型标记区分。 类型标签在之前
输入流中的值并定义字节的大小和种类
随后。
*/

// 文档代码

package rlp
