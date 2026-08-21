<style>
/* ===========================================================
   Academic CV — 参照 moderncv / Awesome-CV 的学术变体
   以及 RenderCV 的 classic 主题

   学术履历天生多页，排版要点跟求职简历相反：
   不压缩、不塞满，靠稳定的层级让评审快速跳到某一节。

   签名细节：小型大写章节标题、编号列表做悬挂缩进（论文列表的
   第二行对齐到第一行的文字，不对齐到序号）、表格当奖项年表用。
   =========================================================== */

:root {
    --font-body: var(--font-serif);

    --fs: 10.2pt;
    --lh: 1.45;
    --fs-name: 19pt;
    --fs-sec: 11.4pt;

    --page-x: 19mm;
    --page-y: 17mm;
    --sp-sec: 5.8mm;
    --sp-item: 3.4mm;
    --sp-line: 1.6mm;

    --ink: #1a1a1a;
    --muted: #4f4f4f;
    --rule: #333;
    --rule-soft: #d8d8d8;
    --tint: #f1f1ef;
}

h1 {
    font-weight: 400;
    font-variant: small-caps;
    letter-spacing: 0.045em;
    margin: 0 0 2mm;
}

h1 + p {
    font-size: calc(var(--fs) * 0.92);
    line-height: 1.55;
    margin: 0 0 var(--sp-sec);
    padding-bottom: 3.2mm;
    border-bottom: 0.5pt solid var(--rule);
}

h2 {
    font-variant: small-caps;
    font-weight: 700;
    letter-spacing: 0.07em;
    padding-bottom: 0.7mm;
    border-bottom: 0.4pt solid var(--rule-soft);
    margin: var(--sp-sec) 0 2mm;
}

h3 { font-weight: 700; }

h4 {
    font-style: italic;
    color: var(--muted);
}

/* --- 论文列表：悬挂缩进 --- */
ol {
    list-style: none;
    counter-reset: pub;
    padding-left: 8mm;
}

ol > li {
    counter-increment: pub;
    position: relative;
    margin-bottom: 1.6mm;
    text-align: justify;
    text-align-last: left;
}

ol > li::before {
    content: "[" counter(pub) "]";
    position: absolute;
    left: -8mm;
    width: 7mm;
    font-variant-numeric: tabular-nums;
    color: var(--muted);
}

ul { padding-left: 5mm; }
li::marker { color: var(--ink); }

/* --- 奖项 / 技能表：年份列窄一点，多页时表头照样不印 --- */
td:first-child {
    width: 18mm;
    font-weight: 400;
    color: var(--muted);
    font-variant-numeric: tabular-nums;
}

td { padding: 0.9mm 0; }

code {
    font-family: var(--font-mono);
    font-size: 0.86em;
    background: none;
    padding: 0;
}

a { text-decoration: underline; text-underline-offset: 1.6pt; text-decoration-thickness: 0.4pt; }

blockquote {
    border: 0;
    margin: 2mm 6mm;
    padding: 0;
    font-size: 0.96em;
}

/* 多页时不要在章节标题后立刻断页 */
h2 { break-after: avoid-page; page-break-after: avoid; }
</style>

# Yiwen Zhao

Ph.D. Candidate in Computer Science, Tsinghua University
Room 4-208, FIT Building, Tsinghua University, Beijing 100084, China
yiwen.zhao@example.edu ⋄ +86 138-0000-0000 ⋄ [yiwenzhao.github.io](https://example.github.io) ⋄ [Google Scholar](https://scholar.google.com)

## Research Interests

Machine learning systems, distributed training efficiency, and memory-aware scheduling for large-scale models. I am particularly interested in making training infrastructure predictable rather than merely fast.

## Education

### Tsinghua University <span>2021.09 – present</span>
#### Ph.D. in Computer Science, advised by Prof. Wei Chen <span>Beijing, China</span>

Dissertation (in progress): *Memory-Aware Scheduling for Distributed Deep Learning*

### Tsinghua University <span>2017.09 – 2021.06</span>
#### B.Eng. in Computer Science and Technology, GPA 3.91/4.0 (rank 4/187) <span>Beijing, China</span>

## Publications

*Underline / bold indicates the author of this CV. † denotes equal contribution.*

1. **Y. Zhao**, L. Sun, W. Chen. *Slack: Memory-Aware Pipeline Scheduling for Large Model Training*. In **OSDI 2025**.
2. **Y. Zhao**†, M. Guo†, W. Chen. *Rethinking Activation Recomputation Under Heterogeneous Memory*. In **ASPLOS 2024**, pp. 331–345.
3. J. Li, **Y. Zhao**, W. Chen. *A Measurement Study of Straggler Effects in Multi-Tenant GPU Clusters*. In **SoCC 2023**, pp. 88–101.
4. **Y. Zhao**, W. Chen. *Towards Predictable Training Throughput*. **HotOS 2023** (workshop).

## Preprints and Under Review

1. **Y. Zhao**, R. Tan, W. Chen. *Elastic Checkpointing Without Global Barriers*. Under review at NSDI 2026. [arXiv:2508.00000](https://arxiv.org)

## Research Experience

### Microsoft Research Asia <span>2024.06 – 2024.12</span>
#### Research Intern, Systems Group, mentor: Dr. Hao Lin <span>Beijing, China</span>

- Designed a scheduler that overlaps checkpoint I/O with backward passes; reduced checkpoint stall time by 71% on a 512-GPU cluster.
- Work contributed to a production training platform and formed the basis of the NSDI 2026 submission.

### Tsinghua Parallel Computing Lab <span>2021.09 – present</span>
#### Graduate Research Assistant <span>Beijing, China</span>

- Built and maintain an open-source profiling toolkit for distributed training (480+ GitHub stars), adopted by three external research groups.
- Led a two-year measurement study of a 2,000-GPU production cluster, resulting in the SoCC 2023 paper.

## Teaching

### Tsinghua University <span>2022 – 2024</span>
#### Teaching Assistant <span>Beijing, China</span>

- *Operating Systems* (undergraduate, ~180 students), Spring 2023, Spring 2024. Redesigned the scheduling lab; median completion rate rose from 62% to 89%.
- *Advanced Computer Architecture* (graduate, ~40 students), Fall 2022.

## Awards and Honors

| Year | Award |
|:--|:--|
| 2025 | ACM SIGOPS Student Travel Grant |
| 2024 | National Scholarship for Graduate Students (top 1%) |
| 2023 | Tsinghua Comprehensive Excellence Award, First Class |
| 2021 | Outstanding Undergraduate Thesis, Tsinghua University |

## Professional Service

- Reviewer / External Reviewer: SOSP 2025, EuroSys 2025, MLSys 2024
- Artifact Evaluation Committee: OSDI 2024, ASPLOS 2023
- Co-organizer, Tsinghua Systems Reading Group (2022 – present)

## Invited Talks

1. *Memory-Aware Scheduling for Large Model Training*. Peking University Systems Seminar, 2025.
2. *What We Learned From Two Years of Cluster Traces*. SoCC 2023, Santa Cruz, CA.

## Skills

| Category | Details |
|:--|:--|
| Programming | C++, Python, CUDA, Go, Rust |
| Systems | PyTorch internals, NCCL, Kubernetes, Slurm, RDMA |
| Languages | Mandarin Chinese (native), English (TOEFL 108, fluent) |

## References

Available upon request.
