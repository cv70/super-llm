# Chat Completions API 改造手册

## 1. 需求概述

本手册将指导如何将现有代码改造为支持 Chat Completions API 的 Gin Web 应用程序。主要功能包括：

- 提供 `/chat/completions` 接口
- 通过 HTTP Header 定义调用的模型列表
- 通过 Header 参数控制响应阶段（一阶段、二阶段、三阶段）
- 保持现有业务逻辑不变

## 2. 当前项目结构分析

根据项目文件结构，我们有以下关键组件：
- `main.go`: 程序入口点
- `domain/committee/`: 委员会相关逻辑
- `pkg/llm/`: 大语言模型相关功能
- `infra/llm.go`: LLM 接口实现
- `config/config.go`: 配置管理

## 3. 改造计划

### 3.1 创建新的 API 路由模块

我们需要创建一个新的包来处理聊天完成请求：

```bash
mkdir -p api/chat
```

### 3.2 实现 Chat Completions Handler

创建 API 处理器来处理 `/chat/completions` 请求。

### 3.3 Header 解析与验证

解析 HTTP Header 中的模型列表和阶段参数。

### 3.4 集成现有 LLM 功能

利用现有的 LLM 实现和委员会逻辑。

## 4. 具体实现步骤

### 步骤 1: 创建 API 包结构

```bash
mkdir -p api/chat
```

### 步骤 2: 创建 Chat Completions Handler

在 `api/chat/handler.go` 文件中实现处理器。

### 步骤 3: 更新主程序以注册路由

修改 `main.go` 来注册新的 API 路由。

### 步骤 4: 创建配置结构

添加必要的配置项用于处理阶段控制。

## 5. API 设计规范

### 5.1 请求格式

```
POST /chat/completions
Content-Type: application/json
X-Models: model1,model2,model3
X-Stage: 1|2|3
```

### 5.2 响应格式

```json
{
  "id": "chatcmpl-1234567890",
  "object": "chat.completion",
  "created": 1677652288,
  "model": "gpt-3.5-turbo",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "Hello! How can I assist you today?"
      },
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 10,
    "completion_tokens": 15,
    "total_tokens": 25
  }
}
```

## 6. Header 参数说明

### 6.1 X-Models
- 格式：逗号分隔的模型名称列表
- 示例：`X-Models: gpt-4,gpt-3.5-turbo`
- 作用：指定要调用的模型列表

### 6.2 X-Stage
- 取值：1 | 2 | 3
- 1：仅返回一阶段结果
- 2：返回一阶段和二阶段结果
- 3：返回所有阶段结果（必须返回）

## 7. 代码结构更新

### 7.1 新增文件
- `api/chat/handler.go` - Chat Completions 处理器
- `api/chat/models.go` - 请求和响应数据结构
- `api/server.go` - API 服务器实现
- `pkg/llm/service.go` - LLM 服务实现

### 7.2 修改文件
- `main.go` - 注册新路由并启动 HTTP 服务器
- `config/config.go` - 添加配置项

## 8. 实现细节

### 8.1 模型选择逻辑
- 从 Header 中提取模型列表
- 验证模型是否存在
- 按顺序调用模型

### 8.2 阶段控制逻辑
- 根据 X-Stage Header 控制返回内容
- 一阶段：基础响应
- 二阶段：增强响应（包含推理过程等）
- 三阶段：完整响应（包含所有信息）

### 8.3 错误处理
- 模型不存在错误
- Header 格式错误
- 内部服务错误