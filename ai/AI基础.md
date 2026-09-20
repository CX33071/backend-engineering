# `AI Agent`开发基础链路

## 一、什么是大模型？

这部分主要是让大家理解大模型是什么。以及它有什么能力和限制。

1.大语言模型，可以理解成一个读过大量数据，能够根据上下文生成文本的超级助理。

大模型的底层工作方式:

```go
输入上下文
    ↓
模型分析
    ↓
预测下一个 Token
    ↓
继续预测
    ↓
生成完整回答
```

2.`token`是什么？

大模型并不是直接理解：我今天心情很好，而是把这句话拆成我/今天/心情/很好，实际`token`怎么且，取决于具体模型。

`token`的2个用途：

(1)`token`决定上下文长度，模型最多支持`N`个`token`，决定的这个上下文长度就是`context window`模型的白板，可以把`context window`理解成：模型一次能够看到的最大信息量。

```go
System Prompt
+
User Prompt
+
历史对话
+
工具结果
+
模型输出
```

这些都要占`token`

(2)`token`和`API`成本有关

`API`会按照输入、输出`token`等维度计费，你的`prompt`越长，历史记录越多，返回结果越多，`token`就越多，花的钱也就越多。

3.`context window`

前面说这就是模型一次能够看到的最大信息量：

```go
Context Window
┌────────────────────────────┐
│ System Prompt              │
│ User Prompt                │
│ 历史对话                    │
│ Tool Call                  │
│ Tool Response              │
│ ...                        │
└────────────────────────────┘
```

所有东西都在这个窗口里，东西会越来越多，所以`agent`做复杂任务的时候，上下文管理非常重要。

4.为什么大模型看起来有记忆

大模型本身并不是通过`API`调用自动记住上一次请求的，第一次`API`请求的时候说我叫小明，第二次`API`请求的什么又不知道自己叫什么了，所以第二次请求的时候如果不传历史的话，模型根本不会知道你之前说过什么。

应用通常会自己维护：

```go
messages = [
    system,
    user,
    assistant,
    user,
    assistant,
    ...
]
```

然后发给模型。所以记忆就是应用层把需要的信息放进`context`。

`system`:给模型规定身份，行为规则，比如：你现在是一个C++工程师

`user`:用户提出任务

`assistant`:模型之前的回答，模型之前说过的话

例如：

```go
[
    {
        "role": "system",
        "content": "你是一个Linux专家"
    },
    {
        "role": "user",
        "content": "解释一下epoll"
    }
]
```

模型看到的是一整个结构化消息。

5.大模型的能力和短板

它擅长：

```go
理解自然语言
      ↓
分析/推理
      ↓
生成内容
      ↓
遵守规则
```

它不擅长：

```go
直接操作外部世界
访问你的数据库
读取你的服务器
获取实时信息
执行程序
长期保存任务状态
```

大模型只有大脑，没有手，不能帮我们去解决问题

## 二、`Prompt`

1.`prompt`是什么？

就是你给模型的输入内容。

可以包括：

```go
Prompt
├── 角色
├── 指令
├── 背景
├── 参考资料
├── 示例
├── 输出格式
└── 约束条件
```

因为模型是在根据上下文来预测接下来该生成什么，如果你上下文模糊的话，答案范围很大，可能的回答很多，结果不稳定。

来举个好的`prompt`:

```go
你是一个技术人员：
请用 3 句话总结下面这份技术文档，
面向非技术管理者，
重点说明业务影响。
```

2.`User Prompt`和`System Prompt`

`User Prompt`就是用户这次具体想让模型做什么，比如：帮我分析这个报错

`System Prompt`就是岗位说明书+行为规范，比如：

```go
你是一名资深 SRE。

规则：
1. 只分析系统故障
2. 不要自行编造背景
3. 信息不足时明确说明
4. 输出必须使用 JSON
```

3.`Prompt`六原则

```go
好的 Prompt
    │
    ├── 角色明确
    ├── 正向约束
    ├── 背景充分
    ├── 输出格式明确
    ├── 给示例
    └── ...
```

最重要的就是背景信息+输出格式

## 三.`Agent`

1.`Agent=LLM+Tools+Memory+Loop`

`Agent`是以大模型为大脑，能够自主感知、决策、调用工具并完成多步骤任务的程序。

大模型只是`agent`的核心组件之一。

2.`agent`的四个核心组件

(1)`LLM`——大脑

负责：

```go
理解任务
分析
推理
决策
```

(2)`Tool`——手

负责：

```go
查数据库
搜索网页
读文件
写文件
发邮件
调用API
执行程序
```

(3)`Memory`——记事本

负责：

```go
短期记忆
    ↓
当前任务上下文

长期记忆
    ↓
外部数据库/存储
```

(4)`Loop`——发动机

```go
思考
 ↓
行动
 ↓
观察
 ↓
再思考
 ↓
行动
 ↓
观察
 ↓
...
```

如果没有`Loop`:

```go
用户 → LLM → 答案
```

这就是普通调用，和大模型聊天就没有`Loop`

有`Loop`:

```go
用户
 ↓
LLM
 ↓
Tool
 ↓
结果
 ↓
LLM
 ↓
Tool
 ↓
结果
 ↓
LLM
 ↓
最终答案
```

这才真正体现`agent`

2.`ReAct`:`agent`最重要的工作方式之一

`ReAct`=`Reasoning+Acting`

```go
Think
 ↓
Act
 ↓
Observe
 ↓
Think
 ↓
Act
 ↓
Observe
 ↓
...
```

`Think`、`Act`、`Observe`这三个流程反复循环，比如现在用户说：帮我分析今晚数据库报警：`agent`:

```go
Think：
我需要先看看报警日志

Act：
调用 get_alert_log()

Observe：
发现大量慢查询

Think：
那我需要查慢查询

Act：
调用 get_slow_query()

Observe：
发现某 SQL 执行时间异常

Think：
再查数据库连接情况

Act：
调用 get_db_status()

...
```

3.`Plan`和`Execute`

如果任务特别复杂：`ReAct`会一步一步想，就会想一步做一步、想一步做一步。容易：

- 迷路
- 绕路
- Token 消耗大
- 上下文越来越长

所以有了：

```go
Planner
   ↓
先规划完整任务
   ↓
Executor
   ↓
一步一步执行
```

`Plan Execute`和`ReAct`的区别就是一个直接上手做，边做边思考，另一个就是先做好计划，做好流程，后面直接跟着计划执行就行。

选择逻辑：

````go
短、灵活的任务
        → ReAct

复杂、多步骤任务
        → Plan and Execute

复杂系统
        → 外层 Plan & Execute
          内层 ReAct
````

4.`agent`和`workflow`的区别

`workflow`就是你提前写死：

```go
A
 ↓
B
 ↓
如果X → C
否则 → D
```

程序决定路径

`agent`:

```go
LLM自己判断
 ↓
调用工具
 ↓
看结果
 ↓
再判断
 ↓
调用工具
```

路径是动态的，`LLM`根据环境决定下一步

## 四.`Tool`

`Tool`就是让大模型能够操作外部世界的能力。

例如：

```go
Tool
├── get_weather()
├── search_web()
├── read_file()
├── write_file()
├── query_mysql()
├── send_email()
└── execute_command()
```

但是`Tool`并不只是调用函数，这个函数存在但是大模型并不知道它的存在，因为模型不会自动扫描你的`python`项目：

```go
Tool
=
函数本体
+
name
+
description
+
parameters
```

1.`Tool`的四个核心部分

(1)函数本体

去执行：

```go
def get_weather(city,date)
```

模型通常不需要知道里面怎么实现

(2)`name`

`get_weather`不要`func_001`，因为模型需要通过名字去理解它的大致作用。

(3)`description`

当用户询问当地天气如何出去是否需要带伞的时候，`agent`会根据`description`判断现在要不要用这个工具，所以`description`写的好工具选择更准确。

(4)`parameters`

告诉`agent`:

```go
这个函数需要什么参数？
参数是什么类型？
哪些是必填？
格式是什么？
```

比如:

```go
{
    "city": {
        "type": "string"
    },
    "date": {
        "type": "string"
    }
}
```

这就是`JSON Schema`，可以理解成这是给模型看的函数参数清单。

2.工具也有风险等级

比如只读的风险等级比写入的风险等级高：

低风险：

```go
查天气
搜索网页
读数据库
读文件
```

中风险：

```go
写数据库
发邮件
发通知
修改配置
```

高风险：

```go
删除数据
执行危险命令
修改生产环境
```

所以真正的`agent`并不能说是`LLM`想干什么就干什么，而是要经过:

```go
LLM决策
 ↓
权限控制
 ↓
风险判断
 ↓
工具执行
```

## 五.`Function Calling`

现在大模型说我要调用`get_weather`，那么`agent`怎么准确知道：

```go
调用哪个函数？
参数是什么？
参数格式是什么？
```

这里就要用到`Function Calling`！

(1)`Function Calling`是什么

`Function Calling`是大模型和`agent`之间标准化的工具调用协议。

(2)`Tool`和`Function Calling`有什么区别

`Tool`=能力

`Function Calling`=调用这个能力的通信机制

```go
Tool
└── get_weather
      ↓
“我有查询天气的能力”


FUntion Call:

LLM
 ↓
我要调用 get_weather
参数：
city = 上海
date = 2026-03-19
 ↓
Agent
 ↓
执行 get_weather()
```

所以`Tool`解决有什么能力，`Function Calling`解决怎么标准化地调用这个能力。

在进行第6部分学习之前我们先来顺一下目前的`Agent`工作流程：

用户现在说帮我查一下明天上海的天气：

**第一步**：`Prompt`，用户提供的“帮我查一下明天上好的天气”这是`User Prompt`

**第二步**：`LLM`模型去理解：

```go
用户想知道：
上海
明天
天气
```

但是模型本身不能直接查询实时天气

**第三步**：`Tool`,`agent`给模型提供了`get_Waether`这个工具

**第四步**：`Function Calling`

模型决定参数给`agent`:

```go
{
    "name": "get_weather",
    "arguments": {
        "city": "上海",
        "date": "2026-09-19"
    }
}
```

这里模型告诉`agent`我要调用这个工具，而且参数是这些

**第五步**：`agent`收到:

```go
get_weather
上海
2026-09-19
```

然后真正去执行：

```go
get_weather("上海", "2026-09-19")
```

**第六步**：`Tool`返回

```go
{
    "temperature": "20~28℃",
    "condition": "晴",
    "wind": "3级"
}
```

**第七步**：结果重新返回给`LLM`

现在模型看到之前的用户问题加之前的上下文加`Tool`回复，然后生成：明天上海天气晴，20，整体适合出行。

## 六、`MCP`

1.为什么需要`MCP`?

`FUnction Calling`解决的是：大模型如何告诉`agnet`我要调用这个工具还有参数。

但是每一个工具都要自己研究`API`，写封装，写`Tool Definition`，写`Function Calling`等等，如果我们现在有10个工具，5个模型，那我们就要写50套代码，`MCP`解决的就是这个问题，`MCP`就是给工具世界建立统一标准。

`MCP`:`Model Context Protocol`，模型上下文协议

以前一个工具一个接口，现在：

```go
工具 A ─┐
工具 B ─┤
工具 C ─┼── MCP
工具 D ─┤
工具 E ─┘
```

`AI`应用只需要支持`MCP`，就可以接入符合`MCP`标准的工具。于是：

```go
原来：

N 个工具 × M 个 AI 应用
        ↓
       N×M


MCP：

N 个 MCP Server
+
M 个 MCP Client
        ↓
       N+M
```

现在只需要每个客户端和服务端对接的这`N+M`套代码就行。

2.`MCP`的三个角色

```go
                Host
                 │
                 ▼
               Client
                 │
        ┌────────┼────────┐
        ▼        ▼        ▼
     Server    Server    Server
     GitHub    MySQL     文件
```

(1)`Host`就是你正在使用的`AI`应用，比如你的`claude`、`codex`。`Host`是整个交互的入口。

(2)`client`是`Host`里的连接器，负责：

```go
找到 Server
建立连接
发送请求
接收结果
```

(3)`server`才是真正提供能力的一方，`Github MCP Server`、`MySQL MCP Server`等它们把自己的能力按照`MCP`标准暴露出来

3.`Function Calling`和`MCP`处于不同层

```go
LLM
 │
 │ Function Calling
 │
 ▼
Agent
 │
 │ MCP
 │
 ▼
MCP Server
 │
 ▼
真实工具
```

`MCP`:`agent`解决去哪里找这个能力，怎么连接它

4.`MCP Server`不只是`Tool`

```go
MCP Server
 ├── Tools
 ├── Resources
 └── Prompts
```

`MCP Server`提供三种能力：

`Tools`执行操作：

```go
发邮件
查数据库
提交代码
搜索网页
```

`Resources`读取数据

`prompts`可复用`prompt`模板：

```go
代码Review模板
日报生成那哦半
故障分析模板
```

## 七、`Skills`

工具有了，但是每次还要告诉`AI`怎么干活，怎么办？

假设你现在每天都让`AI`生成项目日报，你每次都得告诉他：

```go
先查 Jira
↓
再查数据库
↓
再分析昨天完成的任务
↓
按照这个格式生成日报
↓
不要遗漏 Bug
↓
最后输出 Markdown
```

于是有了`skills`，文件把`skills`理解成：给`AI`的`SOP`标准作业流程，也就是将`prompt`+执行步骤+工具+触发条件+上下文这些东西封装成一个可以复用的东西。

1.`prompt`和`skills`的区别

可能会有人问`skill`不就是很多`prompt`吗？

```go
Prompt
=
告诉 AI 怎么做


Skill
=
告诉 AI 怎么做
+
给它准备工具
+
规定工作流程
+
规定什么时候使用
+
把这一套东西保存下来
```

`prompt`是指令，`skill`是可复用的完整工作能力。

2.`skill`文件长什么样？

通常类似：

```go
skills/
└── report-generator/
    └── SKILL.md
```

里面可以包含：

```go
---
name: report-generator
description: 生成项目日报
tools:
  - jira
  - database
---

# 工作流程

1. 获取 Jira 今日任务
2. 查询数据库
3. 分析完成情况
4. 整理 Bug
5. 生成日报
6. 按规定格式输出
```

所以`skill`的核心思想就是把会干这个活从人的经验变成`AI`可以读取和执行的标准流程。也可以说`skill`=操作手册，把红桔+`prompt`+工作流程沉淀成`SOP`

## 八、`RAG`

解决的是：大模型不知道你的私有数据怎么办？

假设你有：

```go
公司运维手册
内部技术文档
历史故障报告
数据库
代码仓库
内部 SOP
```

大模型可能根本没见过，这时你问我们公司的`redis`集群昨天为什么报警，那它肯定不知道呀。这时你想那我把我的东西上传给它不就行了，但是这里不能把所有资料直接塞给`LLM`，主要有3个原因：

(1)`context window`有上限，文档太多，超过上下文窗口

(2)成本高、速度慢

(3)注意力稀释，真正的相关内容可能只有10万字里的500字

那么这里就要用到我们的`RAG`L了。

1.`RAG`:检索增强生成

```go
不要：
全部知识 → Prompt → LLM

而是：

用户问题
   ↓
检索知识库
   ↓
找到最相关的几段
   ↓
放进 Prompt
   ↓
LLM
   ↓
回答
```

先检索，再生成。

2.`RAG`的两个阶段

`RAG`必须掌握的结构：

```go
RAG
├── 离线建库
└── 在线检索生成
```

第一阶段：离线建库

就是把所有数据存到数据库里

`PDF`、`Word`、网页、代码这些东西进入系统：

第一步：加载，把各种数据统一解析成文本

第二步：用户的问题通常只对应其中的一小部分，所以我们`chunking`把长文本切成小块

第三步：`Embedding`，把数据库连接池耗尽变成:

```go
[0.123, -0.532, 0.781, ...]
```

这些就是向量，`Embedding`模型负责：

```go
文本
 ↓
Embedding
 ↓
向量
```

它试图把文字的语义信息编码到向量空间里

第四步：存入向量数据库

```go
Chunk
 +
Embedding Vector
 +
Metadata
        ↓
Vector Database
```

这就完成了离线建库

第二阶段：在线阶段

用户真正提问的时候进入在线阶段：

```go
用户问题
   ↓
Embedding
   ↓
问题向量
   ↓
向量数据库
   ↓
找到最相似 Chunk
   ↓
取回原文
   ↓
拼到 Prompt
   ↓
LLM
   ↓
回答
```

3.`RAG`为什么不是普通搜索？

假设数据库里有："数据库响应时间异常"，但是用户问的是："为什么`SQL`变慢了"，关键词搜索不到，因为字面不一样，但是`Embedding`之后两个向量在语义空间中比较接近，就可以找到。

关键词搜索：找字一样

语义搜索：找意思像

4.`RAG`可以被封装成`agent`的一个`tool`，所以如果以后看到：把`RAG`接入`agrnt`本质上就是：

```go
知识库检索
    ↓
封装成 Tool
    ↓
Agent 调用
```

## 九、向量数据库

这部分就是专门解释为什么`RAG`需要向量数据库的。

为什么不用`MySQL`存向量呢？其实完全可以存，但是`MySQL`不擅长进行大规模向量相似搜索。

1.`MySQL`和向量搜索的区别

`MySQL`擅长精确查询，范围查询，但是向量搜索要求：给我一个向量，找出和它最相似的`Top-K`个向量。这并不是`B+ Tree`擅长的问题

2.最笨的方法：暴力搜索

假设100万个向量，每个向量1536维，暴力搜索的话那这不炸了？！

3.向量数据库真正解决了什么？

在海量向量中快速找到相似向量。

它使用专门的:

```go
ANN
Approximate Nearest Neighbor
近似最近邻搜索
```

为什么是近似？

```go
绝对最优
     ↓
计算特别贵

近似最优
     ↓
计算便宜很多
```

向量数据库不一定保证：我找到数学上绝对最近的那个，而是我很快找到一份非常接近的结果，这就是用极少的精度换大量的速度。

4.`HNSW`

一种非常罕见的`ANN`索引结构，它的核心思想是：

```go
高层
↓
快速跳跃
↓
缩小范围
↓
低层
↓
精细搜索
```

比如说你先要找一个知识点并不是一本书两本书三本书这样去找，而是先找计算机科学相关的书，再找和数据库有关的书，再找索引，看内容

5.几个常见向量数据库

`Chroma`、`Pinecone`、`Milvus`等

## 十、`Harness Engineering`

这部分讨论：怎样让`agent`在真实工程环境里长期稳定地把事情做好

这里有三个比较重要的阶段：

```go
Prompt Engineering
        ↓
Context Engineering
        ↓
Harness Engineering
```

`prompt Engineering`解决让模型听懂

`context Engineering`解决让模型知道该看什么

`Harness Engineering`解决让模型持续做对

1.`prompt Engineering`

核心就是：

```go
角色
+
背景
+
参考资料
+
任务
+
约束
+
输出格式
```

```go
你是资深后端工程师

背景：
这是一个 Go 微服务项目

任务：
分析这个 Bug

约束：
不要修改数据库逻辑

输出：
给出根因 + 修改方案
```

解决的核心问题：怎么让模型理解我的要求？

2.`context Engineering`

`agent`可能要做这些事情:

```go
读 10 个文件
↓
查数据库
↓
看 Git 历史
↓
执行测试
↓
看测试结果
↓
修改代码
↓
重新测试
```

需要解决的问题：模型到底应该在什么时候看到哪些信息，不是信息越多月号，而是在正确的时间把正确的信息放进上下文。

3.`Haeness Engineering`

如果`prompt`正确,`tool`正确但还是一直错怎么办呢，可能你告诉它这个错了，下次注意，但是`agent`不会因为你骂它一次就记住了，于是`Harness Engineering`出现了

`Harness Engineering`:不要只修这一次的错误，要把错误修进环境

(1)修进环境

比如`agent`经常修改代码之后忘记跑测试，你可以增加自动化测试/增加`linter`/增加`Git Hook`，让环境本身约束`agent`，于是：

```go
Agent 犯错
   ↓
找到错误原因
   ↓
修改环境
   ↓
环境变强
   ↓
Agent 下次更不容易犯
```

每次`agent`犯错，都尝试把问题永久沉淀到环境中，从而形成复利。

(2)其实可以理解成给`agent`搭建一个“工程约束+工具+自动检查+工作流程”的运行环境，`agent`不再是模型说什么就是什么，而是：

```go
Agent
  ↓
工程环境
  ├── 规则
  ├── 工具
  ├── 测试
  ├── Linter
  ├── Git Hook
  ├── 文档
  └── 自动验证
```

(3)为什么`Harness Engineering`很重要

模型能力只是整个系统的一部分，真正决定系统能不能稳定工作的还包括：

```go
模型
+
Prompt
+
Context
+
Tools
+
Memory
+
Tests
+
Rules
+
Environment
+
Feedback
```

## 十一、总结

现在把整篇总结起来就是：

```go
                    AI 应用
                       │
                       ▼
                     LLM
                   大模型
                       │
             ┌─────────┴─────────┐
             ▼                   ▼
          Prompt              Context
        怎么告诉它             给它什么
             │                   │
             └─────────┬─────────┘
                       ▼
                     Agent
                       │
             ┌─────────┼─────────┐
             ▼         ▼         ▼
           Tool      Memory      Loop
             │
             ▼
      Function Calling
      “我要调用什么”
             │
             ▼
            MCP
      “工具怎么标准化接入”
             │
             ▼
           Skills
      “这一套活怎么标准化”
             │
             ▼
            RAG
      “怎么访问外部知识”
             │
             ▼
        Embedding
             │
             ▼
      Vector Database
             │
             ▼
       Similarity Search
             │
             ▼
       返回相关知识
             │
             ▼
            LLM
```

再加最后一层：

```go
                Harness Engineering
                        │
        ┌───────────────┼───────────────┐
        ▼               ▼               ▼
      Rules           Tools           Tests
        │               │               │
        └───────────────┼───────────────┘
                        ▼
                稳定运行 Agent
```

这些名词的定义：

```go
LLM:负责理解、推理、生成，是agent的大脑
```

```go
Prompt:告诉模型应该做什么，怎么做以及有什么限制
```

```go
agent:让LLM配合工具，记忆和执行循环完成多步洲任务的系统
```

```go
tool:agrnt可以调用的外部能力
```

```go
function calling:让大模型用结构化的方式告诉agent：调用哪个工具，传什么参数
```

```go
MCP:统一AI应用于外部工具/服务之间连接方式的开放协议
```

```go
Skills:把Prompt,工作流程和工具组合成可复用的标准工作能力/SOP
```

```go
RAG:回答问题之前先检索外部知识，再让模型基于检索结果生成答案
```

```go
向量数据库：专门用于高效进行向量相似度搜索的数据系统，是RAG常见的检索基础设施
```

```go
Harbess Engineering:通过规则、工具、测试、自动化检查和环境设计，让agent在真实工程任务中更加稳定地完成工作
```

这些名词都解决了什么？

```go
问题：
模型只能说话
        ↓
解决：
Tool

问题：
模型怎么准确告诉程序调用哪个工具？
        ↓
解决：
Function Calling

问题：
工具越来越多，接入全靠自己写
        ↓
解决：
MCP

问题：
每次都要重复告诉 AI 怎么完成一项工作
        ↓
解决：
Skills

问题：
模型不知道我的私有知识
        ↓
解决：
RAG

问题：
RAG 需要大量向量检索
        ↓
解决：
Vector Database + ANN/HNSW

问题：
Agent 还是会在复杂工程任务里反复犯错
        ↓
解决：
Harness Engineering
```

`LLM` 是核心模型，`Agent` 是应用架构，`Function Calling` 是调用机制，`MCP `是工具连接标准，`Skills` 是工作流沉淀，`RAG` 是知识增强，向量数据库是 `RAG` 的基础设施，而 `Harness Engineering `则是在更高层解决 `Agent` 的工程稳定性问题。