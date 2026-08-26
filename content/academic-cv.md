<style>
/* ===========================================================
   Academic CV — 参照 moderncv / Awesome-CV 的学术变体
   以及 RenderCV 的 classic 主题

   学术履历天生多页，排版要点跟求职简历相反：
   不压缩、不塞满，靠稳定的层级让评审快速跳到某一节。

   签名细节：小型大写章节标题、编号列表做悬挂缩进（论文列表的
   第二行对齐到第一行的文字，不对齐到序号）、表格当奖项年表用。
   =========================================================== */

body {
    padding: 17mm 19mm;
    font-family: "Latin Modern Roman", XCharter, Charter, "Bitstream Charter",
                 Palatino, Georgia, "Times New Roman", "Source Han Serif SC",
                 "Songti SC", serif;
    font-size: 10.2pt;
    line-height: 1.45;
    color: #1a1a1a;
}

h1 {
    font-weight: 400;
    font-variant: small-caps;
    font-size: 19pt;
    letter-spacing: 0.045em;
    line-height: 1.15;
    margin: 0 0 2mm;
}

h1 + p {
    font-size: 9.38pt;
    line-height: 1.55;
    color: #4f4f4f;
    margin: 0 0 5.8mm;
    padding-bottom: 3.2mm;
    border-bottom: 0.5pt solid #333;
}

h2 {
    font-variant: small-caps;
    font-weight: 700;
    font-size: 11.4pt;
    line-height: 1.25;
    letter-spacing: 0.07em;
    padding-bottom: 0.7mm;
    border-bottom: 0.4pt solid #d8d8d8;
    margin: 5.8mm 0 2mm;
}

h2:first-of-type { margin-top: 0; }

h3 {
    font-weight: 700;
    font-size: 10.2pt;
    line-height: 1.3;
    margin: 3.4mm 0 0;
}

h4 {
    font-style: italic;
    font-weight: 400;
    font-size: 9.38pt;
    line-height: 1.35;
    color: #4f4f4f;
    margin: 0.4mm 0 1.6mm;
}

/* 时间和地点：斜体写在标题末尾，浮到右边 */
h3 em, h4 em { float: right; font-weight: 400; font-variant-numeric: tabular-nums; }
h3 em { font-style: normal; }

p { margin: 0 0 1.6mm; }
ul, ol { margin: 1.6mm 0 2.2mm; padding-left: 4.6mm; }
li { margin: 0 0 0.9mm; }

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
    color: #4f4f4f;
}

ul { padding-left: 5mm; }
li::marker { color: #1a1a1a; font-size: 0.9em; }

/* 行内 code 沿用灰底标签，只是字号收小一点 */
code {
    font-family: inherit;
    font-size: 0.86em;
    padding: 0.3mm 1.4mm;
    border-radius: 0.8mm;
    background: rgba(17, 17, 17, 0.06);
    white-space: nowrap;
}

a { color: inherit; text-decoration: underline; text-underline-offset: 1.6pt; text-decoration-thickness: 0.4pt; }

blockquote {
    border: 0;
    margin: 2mm 6mm;
    padding: 0;
    font-size: 0.96em;
    color: #4f4f4f;
}

/* --- 奖项 / 技能表：年份列窄一点，多页时表头照样不印 --- */
table { width: 100%; border-collapse: collapse; font-size: 0.96em; margin: 1.6mm 0 2.4mm; }
thead { display: none; }
td { padding: 0.9mm 0; vertical-align: top; }

td:first-child {
    width: 1px;
    white-space: nowrap;
    padding-right: 5mm;
    font-weight: 400;
    color: #4f4f4f;
    font-variant-numeric: tabular-nums;
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

### Tsinghua University *2021.09 – present*
#### Ph.D. in Computer Science, advised by Prof. Wei Chen *Beijing, China*

Dissertation (in progress): *Memory-Aware Scheduling for Distributed Deep Learning*

### Tsinghua University *2017.09 – 2021.06*
#### B.Eng. in Computer Science and Technology, GPA 3.91/4.0 (rank 4/187) *Beijing, China*

## Publications

*Underline / bold indicates the author of this CV. † denotes equal contribution.*

1. **Y. Zhao**, L. Sun, W. Chen. *Slack: Memory-Aware Pipeline Scheduling for Large Model Training*. In **OSDI 2025**.
2. **Y. Zhao**†, M. Guo†, W. Chen. *Rethinking Activation Recomputation Under Heterogeneous Memory*. In **ASPLOS 2024**, pp. 331–345.
3. J. Li, **Y. Zhao**, W. Chen. *A Measurement Study of Straggler Effects in Multi-Tenant GPU Clusters*. In **SoCC 2023**, pp. 88–101.
4. **Y. Zhao**, W. Chen. *Towards Predictable Training Throughput*. **HotOS 2023** (workshop).

## Preprints and Under Review

1. **Y. Zhao**, R. Tan, W. Chen. *Elastic Checkpointing Without Global Barriers*. Under review at NSDI 2026. [arXiv:2508.00000](https://arxiv.org)

## Research Experience

### Microsoft Research Asia *2024.06 – 2024.12*
#### Research Intern, Systems Group, mentor: Dr. Hao Lin *Beijing, China*

- Designed a scheduler that overlaps checkpoint I/O with backward passes; reduced checkpoint stall time by 71% on a 512-GPU cluster.
- Work contributed to a production training platform and formed the basis of the NSDI 2026 submission.

### Tsinghua Parallel Computing Lab *2021.09 – present*
#### Graduate Research Assistant *Beijing, China*

- Built and maintain an open-source profiling toolkit for distributed training (480+ GitHub stars), adopted by three external research groups.
- Led a two-year measurement study of a 2,000-GPU production cluster, resulting in the SoCC 2023 paper.

## Teaching

### Tsinghua University *2022 – 2024*
#### Teaching Assistant *Beijing, China*

- *Operating Systems* (undergraduate, ~180 students), Spring 2023, Spring 2024. Redesigned the scheduling lab; median completion rate rose from 62% to 89%.
- *Advanced Computer Architecture* (graduate, ~40 students), Fall 2022.

## Awards and Honors

|  |  |
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

|  |  |
|:--|:--|
| Programming | C++, Python, CUDA, Go, Rust |
| Systems | PyTorch internals, NCCL, Kubernetes, Slurm, RDMA |
| Languages | Mandarin Chinese (native), English (TOEFL 108, fluent) |

## References

Available upon request.
