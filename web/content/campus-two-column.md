<style>
/* ===========================================================
   Deedy — 复刻 github.com/deedy/Deedy-Resume
   （Debarghya Das 那份，学生和应届生里流传最广的双栏）

   签名细节：巨大的居中姓名（姓氏细体）、两栏 4:6 密排、
   章节标题全大写拉开字距后压一条通栏线。
   目标是把课程、项目、经历全塞进一页。

   需要正文里有 <div class="cv-grid"> + .sidebar / .main，
   没有时会自动退化成单栏。
   左栏放教育 / 技能 / 链接 / 课程，右栏放经历 / 项目。

   注意：双栏对 ATS 机器解析不友好，且分页不受 page-break 控制，
         建议只做一页、只投人工渠道。
   =========================================================== */

body {
    padding: 13mm 15mm;
    font-family: Inter, "Source Sans 3", "Helvetica Neue", Helvetica, Arial,
                 "Segoe UI", Roboto, "PingFang SC", "Source Han Sans SC",
                 "Microsoft YaHei", sans-serif;
    font-size: 9.4pt;
    line-height: 1.4;
    color: #1b1b1b;
}

/* --- 通栏头部 --- */
h1 {
    text-align: center;
    font-weight: 300;
    font-size: 27pt;
    letter-spacing: 0.01em;
    line-height: 1.15;
    margin: 0 0 1.6mm;
}

h1 strong { font-weight: 700; }

h1 + p {
    text-align: center;
    font-size: 8.84pt;
    color: #666;
    line-height: 1.5;
    margin: 0 0 5mm;
    padding-bottom: 3.4mm;
    border-bottom: 0.5pt solid #dcdcdc;
}

/* --- 4:6 双栏，中间一条贯穿到底的发丝线 --- */
.cv-grid {
    display: grid;
    grid-template-columns: 36% 1fr;
    column-gap: 8mm;
    align-items: start;
    background: linear-gradient(#dcdcdc, #dcdcdc) no-repeat;
    background-size: 0.4pt 100%;
    background-position: calc(36% + 4mm) 0;
}

.cv-grid > .sidebar { padding-right: 1mm; }
.cv-grid > .main { padding-left: 1mm; }

h2 {
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.16em;
    font-size: 9.8pt;
    line-height: 1.25;
    padding-bottom: 0.8mm;
    border-bottom: 0.7pt solid #1b1b1b;
    margin: 4.6mm 0 1.6mm;
}

/* 两栏各自的头一个章节标题贴着栏顶 */
.sidebar h2:first-of-type,
.main h2:first-of-type { margin-top: 0; }

h3 {
    font-weight: 700;
    font-size: 9.78pt;
    line-height: 1.3;
    margin: 2.8mm 0 0;
}

h4 {
    font-style: italic;
    font-weight: 400;
    font-size: 8.65pt;
    color: #666;
    line-height: 1.35;
    margin: 0.4mm 0 1mm;
}

/* 时间和地点：斜体写在标题末尾，浮到右边 */
h3 em, h4 em {
    float: right;
    font-weight: 400;
    font-size: 0.92em;
    font-variant-numeric: tabular-nums;
}
h3 em { font-style: normal; }

/* 左栏只有 36% 宽，浮动会把标题挤散，在那儿就当普通斜体 */
.sidebar h3 em, .sidebar h4 em { float: none; }

p { margin: 0 0 1mm; }
ul, ol { margin: 1mm 0 1.6mm; padding-left: 4mm; }
li { margin: 0 0 0.6mm; }
li::marker { color: #666; font-size: 0.8em; }
a { color: inherit; text-decoration: none; }

blockquote {
    margin: 1mm 0;
    padding: 0 0 0 3.5mm;
    border-left: 1.2pt solid #1b1b1b;
    color: #666;
}

/* 左栏更紧、更像一张清单 */
.sidebar { font-size: 0.97em; }
.sidebar ul { list-style: none; padding-left: 0; }
.sidebar li { margin-bottom: 0.9mm; }
.sidebar blockquote {
    border: 0;
    padding: 0;
    margin: 0 0 1.2mm;
    font-size: 0.94em;
}

/* 行内 code 当技能标签用：浅底圆角，字体跟正文走 */
code {
    font-family: inherit;
    font-size: 0.94em;
    padding: 0.3mm 1.4mm;
    border-radius: 0.8mm;
    background: #f0f0f0;
    white-space: nowrap;
}
</style>

# 周雨桐

计算机科学与技术 · 2026 届 ⋄ 求职意向：算法工程师 ⋄ 北京 ⋄ 138-0000-0000 ⋄ zhouyutong@example.com

<div class="cv-grid">
<div class="sidebar">

## 教育背景

**北京邮电大学**

> 计算机科学与技术 · 本科
> 2022.09 - 2026.06
> GPA 3.88 / 4.0，排名 5 / 196

## 核心课程

- 数据结构与算法（96）
- 机器学习（95）
- 概率论与数理统计（97）
- 计算机视觉（93）
- 深度学习实践（94）
- 操作系统（91）

## 技能

- Python、C++、Java、SQL
- PyTorch、NumPy、scikit-learn
- Hugging Face、ONNX、TensorRT
- Git、Docker、Linux、Slurm

## 获奖

- 2025 全国大学生数学建模竞赛 国家一等奖
- 2024 ACM-ICPC 亚洲区域赛（南京站）银奖
- 2024 Kaggle 表格赛 Top 3%（38 / 1,420）
- 国家奖学金（2024）
- 校级一等奖学金（2023 - 2025）

## 链接

- [github.com/example](https://github.com/example)
- [example.github.io](https://example.github.io)

## 语言

- 英语 CET-6 598，可读写英文论文

</div>
<div class="main">

## 实习经历

### 某互联网公司 · 搜索算法组 *2025.06 - 2025.09*
#### 算法实习生 *北京*

- 负责查询改写模型的迭代，用对比学习替换原有的规则同义词表，长尾 query 召回率提升 8.3%。
- 把线上精排模型从 FP32 量化到 INT8，推理延迟下降 41%，离线 AUC 仅损失 0.0008。
- 搭建一套特征漂移监控，上线后提前发现两次因上游埋点变更导致的特征失效。

### 某 AI 创业公司 *2024.07 - 2024.09*
#### 算法实习生 *远程*

- 参与多模态检索模块开发，负责图文对数据的清洗与去重流水线，处理 1,200 万条样本。
- 复现并改进一篇 CVPR 论文的关键模块，在内部数据集上 Recall@10 提升 4.7 个点。

## 科研经历

### 北邮模式识别实验室 *2024.03 - 至今*
#### 本科生研究助理 · 导师：李某某 教授 *北京*

- 研究小样本场景下的图像分类，提出一种基于原型校准的方法，在 miniImageNet 上超过基线 2.1 个点。
- 一作论文投稿至 ICASSP 2026，在审。

## 项目经历

### 轻量级向量检索引擎 *2025.02 - 2025.05*
#### 个人项目 · C++ / Python *[github.com/example/vecdb](https://github.com/example/vecdb)*

- 实现 HNSW 索引与磁盘分片，千万级向量下 Recall@10 达 0.97，单机 QPS 3,400。
- 完整的压测报告与 Docker 部署脚本，累计 340 star。

### 课程知识问答助手 *2024.10 - 2024.12*
#### 团队 3 人 · 负责检索与评测 *RAG · 校内比赛一等奖*

- 设计分块与重排策略，人工评测答案可用率从 61% 提升到 84%。
- 建立 200 条题目的评测集，把「感觉变好了」变成可复现的数字。

</div>
</div>
