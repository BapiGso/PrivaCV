<style>
/* ===========================================================
   Neat — 复刻 Typst 生态里的 neat-cv / modern-cv
   （typst.app/universe/package/neat-cv）

   Typst 那批模板的共同气质：无衬线、留白比 LaTeX 舍得给、
   强调色只用在一个很小的地方。这里落在章节标题前的一小段
   实心短杠上，除此之外整页都是黑白。
   =========================================================== */

body {
    padding: 18mm 19mm;
    font-family: Inter, "Source Sans 3", "Helvetica Neue", Helvetica, Arial,
                 "Segoe UI", Roboto, "PingFang SC", "Source Han Sans SC",
                 "Microsoft YaHei", sans-serif;
    font-size: 10pt;
    line-height: 1.56;
    color: #202020;
}

h1 {
    font-size: 22pt;
    font-weight: 600;
    line-height: 1.15;
    letter-spacing: -0.015em;
    margin: 0 0 2mm;
}

h1 + p {
    color: #71716f;
    font-size: 9pt;
    line-height: 1.7;
    margin: 0 0 6.2mm;
}

/* 章节标题前那一小段实心短杠，是这套唯一的着色 */
h2 {
    display: flex;
    align-items: center;
    gap: 2.6mm;
    font-size: 10.2pt;
    font-weight: 600;
    line-height: 1.25;
    text-transform: uppercase;
    letter-spacing: 0.14em;
    margin: 6.2mm 0 2.4mm;
}

h2:first-of-type { margin-top: 0; }

h2::before {
    content: "";
    flex: 0 0 auto;
    width: 4.5mm;
    height: 1.4pt;
    background: #3f5d75;
}

h3 {
    font-size: 10pt;
    font-weight: 600;
    line-height: 1.3;
    letter-spacing: -0.005em;
    margin: 3.8mm 0 0;
}

h4 {
    font-size: 9.2pt;
    font-weight: 400;
    color: #71716f;
    line-height: 1.35;
    letter-spacing: 0.01em;
    margin: 0.4mm 0 1.6mm;
}

/* 时间和地点：斜体写在标题末尾，浮到右边（这套模板里一律用正体） */
h3 em, h4 em {
    float: right;
    font-style: normal;
    font-weight: 400;
    font-size: 0.9em;
    font-variant-numeric: tabular-nums;
    letter-spacing: 0.02em;
    color: #71716f;
}

p { margin: 0 0 1.6mm; }

/* 无点列表：靠缩进和行距分层，不靠符号 */
ul { list-style: none; padding-left: 4.4mm; margin: 1.6mm 0 2.2mm; }

li {
    position: relative;
    margin: 0 0 1.2mm;
}

ul > li::before {
    content: "";
    position: absolute;
    left: -3.4mm;
    top: 0.62em;
    width: 1.3mm;
    height: 1.3mm;
    border-radius: 50%;
    background: #71716f;
}

a { color: #3f5d75; text-decoration: none; }
h1 + p a { color: inherit; }

/* 行内 code 当技能标签用：浅底圆角，字体跟正文走 */
code {
    font-family: inherit;
    font-size: 0.9em;
    padding: 0.3mm 1.4mm;
    border-radius: 1mm;
    background: #f1f2f3;
    white-space: nowrap;
    color: #3b3b3b;
}

blockquote {
    margin: 1.6mm 0;
    padding: 0 0 0 3.5mm;
    border-left: 1.2pt solid #3f5d75;
    color: #71716f;
}

/* GFM 要求表格有表头行，留空即可；这里把它藏起来 */
table { width: 100%; border-collapse: collapse; font-size: 0.96em; margin: 1.6mm 0 2.4mm; }
thead { display: none; }
td { padding: 0.7mm 0; vertical-align: top; }

td:first-child {
    width: 1px;
    white-space: nowrap;
    padding-right: 4mm;
    padding-top: 0.95mm;
    font-weight: 600;
    font-size: 0.92em;
    letter-spacing: 0.02em;
    text-transform: uppercase;
    color: #71716f;
}
</style>

# 林知夏

前端工程师 · 5 年经验 ⋄ 深圳 ⋄ 139-0000-0000 ⋄ linzhixia@example.com ⋄ [linzhixia.dev](https://example.com)

## 技能

|  |  |
|:--|:--|
| 语言 | TypeScript、JavaScript、HTML、CSS |
| 框架 | React、Vue 3、Next.js、Nuxt |
| 工程 | Vite、Webpack、Turborepo、ESLint、Vitest、Playwright |
| 其他 | Node.js、WebGL、Web Vitals 性能优化、可访问性（WCAG 2.1） |

## 工作经历

### 某 SaaS 服务商 *2022.04 - 至今*
#### 前端负责人 *深圳*

- 从零搭建基于 Turborepo 的前端 Monorepo，统一 4 条产品线的构建、Lint 与发布流程，CI 平均耗时从 14 分钟降到 4 分钟。
- 主导设计系统建设：沉淀 62 个组件、完整的 Design Token 与 Storybook 文档，新页面开发工时平均减少 40%。
- 治理首屏性能：路由级代码分割 + 关键 CSS 内联 + 图片 AVIF 化，LCP 从 4.1s 降到 1.3s，移动端跳出率下降 18%。
- 带 4 人前端小组，建立 Code Review 规范与前端周会，线上前端故障从每月 6 起降到 1 起以内。

### 某在线教育公司 *2020.07 - 2022.03*
#### 前端工程师 *广州*

- 负责直播课堂 Web 端，用 Web Worker 承接弹幕解析，主线程长任务占比从 22% 降到 3%，万人同时在线不掉帧。
- 独立完成课件白板的 Canvas 重构，绘制延迟从 120ms 降到 16ms，支持撤销栈与多端同步。
- 推动接入 Sentry + 自建性能上报，白屏问题的平均发现时间从「用户投诉后」变为「发布后 5 分钟内」。

## 项目经历

### 表格可视化编辑器 *2023*
#### 个人开源项目 · React + WebGL *[github.com/example/grid](https://github.com/example/grid)*

- 用 WebGL 渲染替代 DOM 渲染，10 万行数据下滚动稳定 60fps，内存占用比同类库低约 55%。
- 完整的键盘操作与读屏支持，通过 axe-core 全项检查。

## 教育背景

### 中山大学 *2016.09 - 2020.06*
#### 软件工程 · 本科 *GPA 3.6/4.0*

## 其他

- 掘金专栏作者，前端性能系列文章累计阅读 45 万
- 英语 CET-6，可用英文进行技术会议沟通
