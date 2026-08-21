<style>
/* ===========================================================
   Neat — 复刻 Typst 生态里的 neat-cv / modern-cv
   （typst.app/universe/package/neat-cv）

   Typst 那批模板的共同气质：无衬线、留白比 LaTeX 舍得给、
   强调色只用在一个很小的地方。这里落在章节标题前的一小段
   实心短杠上，除此之外整页都是黑白。
   =========================================================== */

:root {
    --font-body: var(--font-sans);

    --accent: #3f5d75;
    --fs: 10pt;
    --lh: 1.56;
    --fs-name: 22pt;
    --fs-sec: 10.2pt;

    --page-x: 19mm;
    --page-y: 18mm;
    --sp-sec: 6.2mm;
    --sp-item: 3.8mm;
    --sp-line: 1.6mm;

    --ink: #202020;
    --muted: #71716f;
    --rule-soft: #e4e4e2;
    --tint: #f1f2f3;
}

h1 {
    font-weight: 600;
    letter-spacing: -0.015em;
    margin: 0 0 2mm;
}

h1 + p {
    font-size: calc(var(--fs) * 0.9);
    line-height: 1.7;
    margin: 0 0 var(--sp-sec);
}

/* 章节标题前那一小段实心短杠，是这套唯一的着色 */
h2 {
    display: flex;
    align-items: center;
    gap: 2.6mm;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.14em;
    margin: var(--sp-sec) 0 2.4mm;
}

h2::before {
    content: "";
    flex: 0 0 auto;
    width: 4.5mm;
    height: 1.4pt;
    background: var(--accent);
}

h3 {
    font-weight: 600;
    letter-spacing: -0.005em;
}

h4 {
    color: var(--muted);
    letter-spacing: 0.01em;
}

[data-split] > .x {
    color: var(--muted);
    font-size: 0.9em;
    letter-spacing: 0.02em;
}

/* 无点列表：靠缩进和行距分层，不靠符号 */
ul { list-style: none; padding-left: 4.4mm; }

li {
    position: relative;
    margin-bottom: 1.2mm;
}

ul > li::before {
    content: "";
    position: absolute;
    left: -3.4mm;
    top: 0.62em;
    width: 1.3mm;
    height: 1.3mm;
    border-radius: 50%;
    background: var(--muted);
}

a { color: var(--accent); }
h1 + p a { color: inherit; }

code {
    border-radius: 1mm;
    font-size: 0.9em;
    color: #3b3b3b;
}

blockquote {
    border-left-color: var(--accent);
    color: var(--muted);
}

td:first-child {
    font-weight: 600;
    letter-spacing: 0.02em;
    color: var(--muted);
    font-size: 0.92em;
    text-transform: uppercase;
    padding-top: 0.95mm;
}
</style>

# 林知夏

前端工程师 · 5 年经验 ⋄ 深圳 ⋄ 139-0000-0000 ⋄ linzhixia@example.com ⋄ [linzhixia.dev](https://example.com)

## 技能

| 类别 | 内容 |
|:--|:--|
| 语言 | TypeScript、JavaScript、HTML、CSS |
| 框架 | React、Vue 3、Next.js、Nuxt |
| 工程 | Vite、Webpack、Turborepo、ESLint、Vitest、Playwright |
| 其他 | Node.js、WebGL、Web Vitals 性能优化、可访问性（WCAG 2.1） |

## 工作经历

### 某 SaaS 服务商 <span>2022.04 - 至今</span>
#### 前端负责人 <span>深圳</span>

- 从零搭建基于 Turborepo 的前端 Monorepo，统一 4 条产品线的构建、Lint 与发布流程，CI 平均耗时从 14 分钟降到 4 分钟。
- 主导设计系统建设：沉淀 62 个组件、完整的 Design Token 与 Storybook 文档，新页面开发工时平均减少 40%。
- 治理首屏性能：路由级代码分割 + 关键 CSS 内联 + 图片 AVIF 化，LCP 从 4.1s 降到 1.3s，移动端跳出率下降 18%。
- 带 4 人前端小组，建立 Code Review 规范与前端周会，线上前端故障从每月 6 起降到 1 起以内。

### 某在线教育公司 <span>2020.07 - 2022.03</span>
#### 前端工程师 <span>广州</span>

- 负责直播课堂 Web 端，用 Web Worker 承接弹幕解析，主线程长任务占比从 22% 降到 3%，万人同时在线不掉帧。
- 独立完成课件白板的 Canvas 重构，绘制延迟从 120ms 降到 16ms，支持撤销栈与多端同步。
- 推动接入 Sentry + 自建性能上报，白屏问题的平均发现时间从「用户投诉后」变为「发布后 5 分钟内」。

## 项目经历

### 表格可视化编辑器 <span>2023</span>
#### 个人开源项目 · React + WebGL <span>[github.com/example/grid](https://github.com/example/grid)</span>

- 用 WebGL 渲染替代 DOM 渲染，10 万行数据下滚动稳定 60fps，内存占用比同类库低约 55%。
- 完整的键盘操作与读屏支持，通过 axe-core 全项检查。

## 教育背景

### 中山大学 <span>2016.09 - 2020.06</span>
#### 软件工程 · 本科 <span>GPA 3.6/4.0</span>

## 其他

- 掘金专栏作者，前端性能系列文章累计阅读 45 万
- 英语 CET-6，可用英文进行技术会议沟通
