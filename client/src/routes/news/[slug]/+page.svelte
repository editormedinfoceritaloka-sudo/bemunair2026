<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/state';

  import type { JSONContent } from '@tiptap/core';

  import RichTextRenderer from '$lib/components/news/RichTextRenderer.svelte';

  type BlogArticle = {
    id: string;
    slug: string;
    title: string;
    description: string;
    cover: string;
    category: string;
    author: {
      name: string;
      role: string;
      avatar: string;
    };
    publishedAt: string;
    publishedAtISO: string;
    readTime: string;
    content: JSONContent;
  };

  const article: BlogArticle = {
    id: 'sop-bem-unair-2026',

    slug:
      'penyusunan-standar-operasional-procedure',

    title:
      'Penyusunan Standar Operasional Procedure',

    description:
      'Langkah strategis untuk menjaga standardisasi, koordinasi, dan keberlangsungan tata kelola organisasi BEM UNAIR 2026.',

    cover:
      '/program-kerja/test-1.png',

    category:
      'Program Kerja',

    author: {
      name:
        'Kementerian Pendayagunaan Aparatur Kabinet',

      role:
        'BEM Universitas Airlangga 2026',

      avatar:
        '/logo/logo-kabinet.png'
    },

    publishedAt:
      '4 Agustus 2026',

    publishedAtISO:
      '2026-08-04',

    readTime:
      '5 menit baca',

    content: {
      type: 'doc',

      content: [
        {
          type: 'paragraph',

          content: [
            {
              type: 'text',

              text:
                'Setiap kementerian atau divisi memiliki pola kerja, tantangan, serta tanggung jawab yang berbeda. '
            },

            {
              type: 'text',

              marks: [
                {
                  type: 'bold'
                }
              ],

              text:
                'Tanpa standardisasi yang jelas'
            },

            {
              type: 'text',

              text:
                ', ego sektoral dapat tumbuh dan koordinasi antarbagian berpotensi terhambat karena adanya perbedaan pemahaman terhadap alur kerja organisasi.'
            }
          ]
        },

        {
          type: 'heading',

          attrs: {
            level: 2
          },

          content: [
            {
              type: 'text',

              text:
                'Mengapa SOP Dibutuhkan?'
            }
          ]
        },

        {
          type: 'paragraph',

          content: [
            {
              type: 'text',

              text:
                'Standar Operasional Procedure menjadi pedoman bersama yang membantu seluruh fungsionaris memahami tahapan kerja, pembagian tanggung jawab, sistem koordinasi, serta mekanisme evaluasi yang berlaku selama satu periode kepengurusan.'
            }
          ]
        },

        {
          type: 'paragraph',

          content: [
            {
              type: 'text',

              text:
                'Melalui dokumen yang tersusun dengan baik, setiap kementerian dapat mengetahui batas kewenangan dan hubungan kerjanya dengan kementerian lain. Hal ini membantu organisasi bergerak lebih konsisten dan mengurangi kesalahan komunikasi.'
            }
          ]
        },

        {
          type: 'blockquote',

          content: [
            {
              type: 'paragraph',

              content: [
                {
                  type: 'text',

                  text:
                    'SOP bukan sekadar dokumen administratif, tetapi fondasi agar organisasi dapat bergerak secara terarah, konsisten, dan berkelanjutan.'
                }
              ]
            }
          ]
        },

        {
          type: 'image',

          attrs: {
            src:
              '/program-kerja/test-2.png',

            alt:
              'Dokumentasi penyusunan SOP BEM UNAIR 2026',

            title:
              'Penyusunan SOP BEM UNAIR 2026'
          }
        },

        {
          type: 'heading',

          attrs: {
            level: 2
          },

          content: [
            {
              type: 'text',

              text:
                'Tujuan Penyusunan SOP'
            }
          ]
        },

        {
          type: 'paragraph',

          content: [
            {
              type: 'text',

              text:
                'Penyusunan SOP BEM UNAIR 2026 dilakukan untuk membangun tata kelola organisasi yang lebih terstruktur. Beberapa tujuan utama yang ingin dicapai meliputi:'
            }
          ]
        },

        {
          type: 'bulletList',

          content: [
            {
              type: 'listItem',

              content: [
                {
                  type: 'paragraph',

                  content: [
                    {
                      type: 'text',

                      text:
                        'Menyamakan pemahaman mengenai alur kerja antarkementerian.'
                    }
                  ]
                }
              ]
            },

            {
              type: 'listItem',

              content: [
                {
                  type: 'paragraph',

                  content: [
                    {
                      type: 'text',

                      text:
                        'Mengurangi kesalahan akibat perbedaan prosedur dan komunikasi.'
                    }
                  ]
                }
              ]
            },

            {
              type: 'listItem',

              content: [
                {
                  type: 'paragraph',

                  content: [
                    {
                      type: 'text',

                      text:
                        'Memastikan setiap pekerjaan memiliki penanggung jawab yang jelas.'
                    }
                  ]
                }
              ]
            },

            {
              type: 'listItem',

              content: [
                {
                  type: 'paragraph',

                  content: [
                    {
                      type: 'text',

                      text:
                        'Membangun sistem organisasi yang dapat dipantau dan dievaluasi.'
                    }
                  ]
                }
              ]
            },

            {
              type: 'listItem',

              content: [
                {
                  type: 'paragraph',

                  content: [
                    {
                      type: 'text',

                      text:
                        'Menjaga kesinambungan pengetahuan organisasi untuk kepengurusan berikutnya.'
                    }
                  ]
                }
              ]
            }
          ]
        },

        {
          type: 'heading',

          attrs: {
            level: 2
          },

          content: [
            {
              type: 'text',

              text:
                'Tahapan Pelaksanaan'
            }
          ]
        },

        {
          type: 'paragraph',

          content: [
            {
              type: 'text',

              text:
                'Penyusunan SOP dilaksanakan secara bertahap agar dokumen yang dihasilkan benar-benar sesuai dengan kebutuhan setiap unit kerja.'
            }
          ]
        },

        {
          type: 'orderedList',

          attrs: {
            start: 1
          },

          content: [
            {
              type: 'listItem',

              content: [
                {
                  type: 'paragraph',

                  content: [
                    {
                      type: 'text',

                      marks: [
                        {
                          type: 'bold'
                        }
                      ],

                      text:
                        'Pemetaan kebutuhan.'
                    },

                    {
                      type: 'text',

                      text:
                        ' Setiap kementerian menyampaikan pola kerja, kebutuhan prosedural, serta kendala yang pernah dihadapi.'
                    }
                  ]
                }
              ]
            },

            {
              type: 'listItem',

              content: [
                {
                  type: 'paragraph',

                  content: [
                    {
                      type: 'text',

                      marks: [
                        {
                          type: 'bold'
                        }
                      ],

                      text:
                        'Penyusunan rancangan.'
                    },

                    {
                      type: 'text',

                      text:
                        ' Tim menyusun rancangan SOP berdasarkan hasil pemetaan dan pembagian kewenangan.'
                    }
                  ]
                }
              ]
            },

            {
              type: 'listItem',

              content: [
                {
                  type: 'paragraph',

                  content: [
                    {
                      type: 'text',

                      marks: [
                        {
                          type: 'bold'
                        }
                      ],

                      text:
                        'Review dan sinkronisasi.'
                    },

                    {
                      type: 'text',

                      text:
                        ' Rancangan diperiksa bersama untuk menghindari tumpang tindih prosedur antarkementerian.'
                    }
                  ]
                }
              ]
            },

            {
              type: 'listItem',

              content: [
                {
                  type: 'paragraph',

                  content: [
                    {
                      type: 'text',

                      marks: [
                        {
                          type: 'bold'
                        }
                      ],

                      text:
                        'Pengesahan dan penerapan.'
                    },

                    {
                      type: 'text',

                      text:
                        ' Dokumen final disahkan dan digunakan sebagai pedoman kerja seluruh kabinet.'
                    }
                  ]
                }
              ]
            },

            {
              type: 'listItem',

              content: [
                {
                  type: 'paragraph',

                  content: [
                    {
                      type: 'text',

                      marks: [
                        {
                          type: 'bold'
                        }
                      ],

                      text:
                        'Monitoring dan evaluasi.'
                    },

                    {
                      type: 'text',

                      text:
                        ' Penerapan SOP ditinjau secara berkala untuk memastikan prosedurnya tetap relevan.'
                    }
                  ]
                }
              ]
            }
          ]
        },

        {
          type: 'heading',

          attrs: {
            level: 3
          },

          content: [
            {
              type: 'text',

              text:
                'Prinsip yang Digunakan'
            }
          ]
        },

        {
          type: 'paragraph',

          content: [
            {
              type: 'text',

              text:
                'Setiap prosedur disusun dengan memperhatikan prinsip efektivitas, akuntabilitas, keterukuran, serta kemudahan penerapan. SOP tidak dibuat untuk memperpanjang birokrasi, melainkan untuk membantu pekerjaan berlangsung lebih cepat dan jelas.'
            }
          ]
        },

        {
          type: 'horizontalRule'
        },

        {
          type: 'heading',

          attrs: {
            level: 2
          },

          content: [
            {
              type: 'text',

              text:
                'Menjaga Keberlanjutan Organisasi'
            }
          ]
        },

        {
          type: 'paragraph',

          content: [
            {
              type: 'text',

              text:
                'Penyusunan SOP diharapkan mampu menciptakan tata kelola organisasi yang lebih transparan, terukur, dan adaptif. Melalui pedoman yang sama, seluruh kementerian dapat bergerak sebagai satu kesatuan dalam mewujudkan tujuan Kabinet Cerita Loka.'
            }
          ]
        },

        {
          type: 'paragraph',

          content: [
            {
              type: 'text',

              text:
                'Dokumen ini juga menjadi bentuk penyimpanan pengetahuan organisasi. Pengalaman, pembelajaran, dan sistem kerja yang telah dibangun tidak berhenti pada satu periode, tetapi dapat menjadi pijakan bagi kepengurusan selanjutnya.'
            }
          ]
        },

        {
          type: 'paragraph',

          content: [
            {
              type: 'text',

              marks: [
                {
                  type: 'italic'
                }
              ],

              text:
                'Setiap langkah kecil yang dilakukan bersama akan menjadi bagian penting dari perjalanan besar BEM UNAIR 2026.'
            }
          ]
        }
      ]
    }
  };

  const routeSlug = $derived(
    page.params.slug
  );

  const titleLines = $derived.by(() => {
    const words = article.title
      .trim()
      .split(/\s+/)
      .filter(Boolean);

    if (words.length === 0) {
      return {
        first: '',
        second: ''
      };
    }

    if (words.length === 1) {
      return {
        first: words[0],
        second: ''
      };
    }

    const splitIndex = Math.floor(
      words.length / 2
    );

    return {
      first: words
        .slice(0, splitIndex)
        .join(' '),

      second: words
        .slice(splitIndex)
        .join(' ')
    };
  });

  onMount(() => {
    const previousScrollRestoration =
      window.history.scrollRestoration;

    const previousScrollBehavior =
      document.documentElement
        .style.scrollBehavior;

    window.history.scrollRestoration =
      'manual';

    document.documentElement
      .style.scrollBehavior = 'auto';

    function scrollToTop(): void {
      window.scrollTo({
        top: 0,
        left: 0,
        behavior: 'auto'
      });

      document.documentElement.scrollTop =
        0;

      document.body.scrollTop = 0;
    }

    scrollToTop();

    const frame = requestAnimationFrame(
      () => {
        scrollToTop();

        document.documentElement
          .style.scrollBehavior =
            previousScrollBehavior;
      }
    );

    return () => {
      cancelAnimationFrame(frame);

      window.history.scrollRestoration =
        previousScrollRestoration;

      document.documentElement
        .style.scrollBehavior =
          previousScrollBehavior;
    };
  });
</script>

<svelte:head>
  <title>{article.title}</title>

  <meta
    name="description"
    content={article.description}
  />

  <meta
    property="og:type"
    content="article"
  />

  <meta
    property="og:title"
    content={article.title}
  />

  <meta
    property="og:description"
    content={article.description}
  />

  <meta
    property="og:image"
    content={article.cover}
  />

  <meta
    name="article-slug"
    content={routeSlug}
  />
</svelte:head>

<article
  class="
    min-h-screen overflow-hidden
    bg-blue-50
    pb-28 pt-24
    sm:pt-28
    lg:pb-36
    lg:pt-36
  "
>
  <header
    class="
      mx-auto flex w-full
      max-w-7xl flex-col
      items-center px-5
      text-center
      sm:px-8
      lg:px-12
    "
  >
    <h1
      class="
        text-[clamp(2.3rem,6.4vw,6rem)]
        leading-[0.9] font-black
        tracking-[-0.065em]
      "
    >
      <span
        class="
          block text-blue-300
          [-webkit-text-stroke:2px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_9px_0_rgba(30,64,175,0.5)]
          sm:[-webkit-text-stroke:3px_white]
        "
      >
        {titleLines.first}
      </span>

      {#if titleLines.second}
        <span
          class="
            relative left-[0.35em]
            block text-blue-700
            [-webkit-text-stroke:2px_white]
            [paint-order:stroke_fill]
            drop-shadow-[0_9px_0_rgba(5,34,70,0.68)]
            sm:[-webkit-text-stroke:3px_white]
          "
        >
          {titleLines.second}
        </span>
      {/if}
    </h1>

    <p
      class="
        mt-9 max-w-3xl
        text-base leading-7
        font-semibold
        text-blue-950/65
        sm:text-lg sm:leading-8
      "
    >
      {article.description}
    </p>

    <div
      class="
        mt-8 flex flex-wrap
        items-center justify-center
        gap-x-4 gap-y-4
      "
    >
      <div
        class="
          flex items-center
          gap-3
        "
      >
        <img
          src={article.author.avatar}
          alt={article.author.name}
          class="
            size-12 rounded-full
            border-2 border-white
            bg-blue-900
            object-cover
            shadow-md
          "
        />

        <div class="text-left">
          <p
            class="
              text-sm font-black
              text-blue-950
            "
          >
            {article.author.name}
          </p>

          <p
            class="
              text-xs font-semibold
              text-blue-950/55
            "
          >
            {article.author.role}
          </p>
        </div>
      </div>

      <span
        aria-hidden="true"
        class="
          hidden h-9 w-px
          bg-blue-900/15
          sm:block
        "
      ></span>

      <p
        class="
          text-xs font-black
          text-blue-700
          sm:text-sm
        "
      >
        {article.category}
      </p>

      <span
        aria-hidden="true"
        class="
          hidden h-9 w-px
          bg-blue-900/15
          sm:block
        "
      ></span>

      <div
        class="
          flex items-center gap-2
          text-xs font-bold
          text-blue-950/55
          sm:text-sm
        "
      >
        <time
          datetime={article.publishedAtISO}
        >
          {article.publishedAt}
        </time>

        <span aria-hidden="true">
          •
        </span>

        <span>
          {article.readTime}
        </span>
      </div>
    </div>
  </header>

  <div
    class="
      mx-auto mt-14
      w-full max-w-7xl
      px-5
      sm:px-8
      lg:px-12
    "
  >
    <figure
      class="
        relative overflow-hidden
        rounded-[28px]
        border-[6px] border-white
        bg-blue-950
        shadow-[0_24px_60px_rgba(5,34,70,0.24)]
      "
    >
      <div
        aria-hidden="true"
        class="
          pointer-events-none
          absolute inset-0 z-10
          bg-gradient-to-t
          from-blue-950/35
          via-transparent
          to-transparent
        "
      ></div>

      <img
        src={article.cover}
        alt={`Cover ${article.title}`}
        class="
          aspect-[16/8]
          w-full object-cover
          object-center
        "
      />
    </figure>
  </div>

  <main
    class="
      relative mx-auto mt-16
      w-full max-w-7xl
      px-5
      sm:px-8
    "
  >
    <div
      aria-hidden="true"
      class="
        pointer-events-none absolute
        -left-20 top-40
        size-72 rounded-full
        bg-blue-300/15
        blur-[100px]
      "
    ></div>

    <RichTextRenderer
      content={article.content}
    />
  </main>

  <footer
    class="
      mx-auto mt-14
      flex w-full max-w-4xl
      flex-col gap-6
      px-5
      sm:px-8
    "
  >
    <div
      class="
        h-px w-full
        bg-gradient-to-r
        from-transparent
        via-blue-900/20
        to-transparent
      "
    ></div>

    <div
      class="
        flex flex-col items-center
        justify-between gap-5
        text-center
        sm:flex-row
        sm:text-left
      "
    >
      <div>
        <p
          class="
            text-sm font-black
            text-blue-950
          "
        >
          Bagikan cerita ini
        </p>

        <p
          class="
            mt-1 text-xs
            font-medium
            text-blue-950/50
          "
        >
          Bersama merajut cerita untuk Airlangga dan Indonesia.
        </p>
      </div>

      <a
        href="/news"
        class="
          inline-flex items-center
          justify-center
          rounded-full
          bg-blue-50 backdrop-blur-sm
          px-6 py-3
          text-sm 
          text-white
          shadow-[0_7px_0_rgba(30,64,175,0.28)]
          transition
          hover:-translate-y-1
          hover:bg-blue-100
          active:translate-y-1
          active:shadow-none
        "
      >
        Kembali ke Berita
      </a>
    </div>
  </footer>
</article>