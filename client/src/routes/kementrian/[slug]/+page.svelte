<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  import MemberCard from '$lib/components/kementrian/MemberCard.svelte';
  import ListProker from '$lib/components/kementrian/ListProker.svelte';

  type DescriptionPart = {
    text: string;
    highlighted: boolean;
  };

  type Member = {
    id: string;
    role: string;
    title: string;
    image: string;
    featured?: boolean;
  };

  type ProgramKerja = {
    id: string;
    title: string;
    description: string;
    slug: string;
  };

  let title = $state(
    'Pendayagunaan Aparatur Kabinet'
  );

  let description = $state(
    'Pendayagunaan Aparatur Kabinet (PAK) merupakan salah satu unit strategis dalam Pengurus Inti Badan Eksekutif Mahasiswa (BEM) Universitas Airlangga (UNAIR) 2026 yang berperan mengelola internal organisasi.'
  );

  const members: Member[] = [
    {
      id: 'dirjen-tata-laksana',
      role: 'Dirjen',
      title: 'Tata Laksana Kerja',
      image: '/kementrian/test.png'
    },
    {
      id: 'menteri-pak',
      role: 'Menteri',
      title: 'Pendayagunaan Aparatur Kabinet',
      image: '/kementrian/test.png',
      featured: true
    },
    {
      id: 'dirjen-audit',
      role: 'Dirjen',
      title: 'Audit dan Penjaminan Mutu',
      image: '/kementrian/test.png'
    }
  ];

  const programs: ProgramKerja[] = [
    {
      id: 'penyusunan-sop',
      title: 'Penyusunan Standar Operasional Procedure',
      description:
        'Penyusunan Standar Operasional Procedure merupakan program kerja yang bergerak pada bidang penyusunan standardisasi kinerja yang harus diikuti oleh seluruh fungsionaris BEM UNAIR 2026.',
      slug: 'penyusunan-standar-operasional-procedure'
    },
    {
      id: 'hearing-kementerian',
      title: 'Hearing Kementerian',
      description:
        'Hearing Kementerian merupakan program kerja yang dilakukan dalam skala satu bulan sekali untuk membahas progres, kendala, hambatan, dan evaluasi terhadap program kerja kementerian.',
      slug: 'hearing-kementerian'
    },
    {
      id: 'monitoring-evaluasi',
      title: 'Monitoring Evaluasi',
      description:
        'Monitoring evaluasi merupakan program kerja yang dilaksanakan tiga bulan sekali dalam satu periode untuk membahas progres ataupun kendala pada setiap kementerian BEM UNAIR 2026.',
     slug: 'monitoring-evaluasi'
    },
    {
      id: 'audit-internal',
      title: 'Audit Internal Kabinet',
      description:
        'Audit Internal Kabinet merupakan kegiatan peninjauan tata kelola organisasi guna memastikan seluruh kementerian bekerja sesuai standar, target, dan nilai Kabinet Cerita Loka.',
        slug: 'audit-internal-kabinet'
    },
    {
      id: 'rapat-koordinasi-kabinet',
      title: 'Rapat Koordinasi Kabinet',
      description:
        'Rapat Koordinasi Kabinet menjadi ruang penyelarasan arah kerja, target, dan kebutuhan lintas kementerian agar pelaksanaan program kabinet tetap terintegrasi.',
        slug: 'rapat-koordinasi-kabinet'
    },
    {
      id: 'evaluasi-tengah-periode',
      title: 'Evaluasi Tengah Periode',
      description:
        'Evaluasi Tengah Periode dilakukan untuk mengukur pencapaian setiap kementerian sekaligus menentukan perbaikan yang diperlukan pada paruh berikutnya.',
        slug: 'evaluasi-tengah-periode'
    },
    {
      id: 'database-fungsionaris',
      title: 'Database Fungsionaris',
      description:
        'Database Fungsionaris mengelola data internal anggota kabinet secara terstruktur untuk mendukung administrasi, pemetaan sumber daya, dan kebutuhan organisasi.',
        slug: 'database-fungsionaris'
    },
    {
      id: 'pemetaan-kinerja-kementerian',
      title: 'Pemetaan Kinerja Kementerian',
      description:
        'Pemetaan Kinerja Kementerian menyajikan gambaran capaian, hambatan, dan kebutuhan pengembangan dari masing-masing unit kerja dalam kabinet.',
        slug: 'pemetaan-kinerja-kementerian'
    },
    {
      id: 'forum-sekretaris-kementerian',
      title: 'Forum Sekretaris Kementerian',
      description:
        'Forum Sekretaris Kementerian menjadi ruang koordinasi administrasi untuk menyelaraskan dokumen, pelaporan, arsip, dan tata kerja seluruh kementerian.',
        slug: 'forum-sekretaris-kementerian'
    },
    {
      id: 'penyelarasan-administrasi',
      title: 'Penyelarasan Administrasi',
      description:
        'Penyelarasan Administrasi memastikan format dokumen dan alur persetujuan internal diterapkan secara konsisten oleh seluruh unit organisasi.',
        slug: 'penyelarasan-administrasi'
    },
    {
      id: 'review-indikator-kinerja',
      title: 'Review Indikator Kinerja',
      description:
        'Review Indikator Kinerja dilakukan untuk memastikan setiap target bersifat terukur, relevan, dan sesuai dengan arah strategis Kabinet Cerita Loka.',
        slug: 'review-indikator-kinerja'
    },
    {
      id: 'sistem-pelaporan-progres',
      title: 'Sistem Pelaporan Progres',
      description:
        'Sistem Pelaporan Progres menyediakan mekanisme pelaporan berkala agar perkembangan program kerja dapat dipantau dan dievaluasi dengan lebih efektif.',
        slug: 'sistem-pelaporan-progres'
    },
    {
      id: 'pendampingan-tata-kelola',
      title: 'Pendampingan Tata Kelola',
      description:
        'Pendampingan Tata Kelola membantu kementerian menyelesaikan persoalan prosedural, pembagian kerja, administrasi, dan koordinasi internal.',
        slug: 'pendampingan-tata-kelola'
    },
    {
      id: 'apresiasi-fungsionaris',
      title: 'Apresiasi Fungsionaris',
      description:
        'Apresiasi Fungsionaris memberikan penghargaan terhadap kontribusi anggota kabinet sebagai upaya membangun budaya organisasi yang sehat dan suportif.',
        slug: 'apresiasi-fungsionaris'
    },
    {
      id: 'evaluasi-akhir-periode',
      title: 'Evaluasi Akhir Periode',
      description:
        'Evaluasi Akhir Periode merangkum keberhasilan, hambatan, dan pembelajaran organisasi sebagai dasar perbaikan untuk kepengurusan berikutnya.',
        slug: 'evaluasi-akhir-periode'
    },
    {
      id: 'laporan-kinerja-kabinet',
      title: 'Laporan Kinerja Kabinet',
      description:
        'Laporan Kinerja Kabinet menyatukan hasil pelaksanaan program, capaian indikator, serta rekomendasi pengembangan organisasi dalam satu dokumen.',
        slug: 'laporan-kinerja-kabinet'
    }
  ];

  let sectionElement!: HTMLElement;
  let glowElement!: HTMLDivElement;
  let decorationElement!: HTMLDivElement;

  let firstTitleElement!: HTMLSpanElement;
  let secondTitleElement: HTMLSpanElement | undefined;
  let descriptionElement!: HTMLParagraphElement;

  let membersSectionElement!: HTMLDivElement;
  let membersHeadingElement!: HTMLHeadingElement;
  let membersGridElement!: HTMLDivElement;

  let programSectionElement!: HTMLDivElement;
  let programHeadingElement!: HTMLHeadingElement;
  let programContentElement!: HTMLDivElement;

  const titleLines = $derived.by(() => {
    const words = title
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

  const descriptionParts =
    $derived.by<DescriptionPart[]>(() => {
      const normalizedTitle = title.trim();

      if (!normalizedTitle) {
        return [
          {
            text: description,
            highlighted: false
          }
        ];
      }

      const lowerDescription =
        description.toLocaleLowerCase('id-ID');

      const lowerTitle =
        normalizedTitle.toLocaleLowerCase('id-ID');

      const parts: DescriptionPart[] = [];
      let currentIndex = 0;

      while (currentIndex < description.length) {
        const matchIndex =
          lowerDescription.indexOf(
            lowerTitle,
            currentIndex
          );

        if (matchIndex === -1) {
          parts.push({
            text: description.slice(currentIndex),
            highlighted: false
          });

          break;
        }

        if (matchIndex > currentIndex) {
          parts.push({
            text: description.slice(
              currentIndex,
              matchIndex
            ),
            highlighted: false
          });
        }

        parts.push({
          text: description.slice(
            matchIndex,
            matchIndex + normalizedTitle.length
          ),
          highlighted: true
        });

        currentIndex =
          matchIndex + normalizedTitle.length;
      }

      return parts;
    });

  onMount(() => {
    gsap.registerPlugin(ScrollTrigger);

    const previousScrollRestoration =
      window.history.scrollRestoration;

    const previousScrollBehavior =
      document.documentElement.style.scrollBehavior;

    let context: gsap.Context | undefined;
    let secondFrame = 0;

    window.history.scrollRestoration = 'manual';
    document.documentElement.style.scrollBehavior = 'auto';

    function forceScrollToTop(): void {
      window.scrollTo({
        top: 0,
        left: 0,
        behavior: 'auto'
      });

      document.documentElement.scrollTop = 0;
      document.body.scrollTop = 0;
    }

    function setupAnimations(): void {
      const reduceMotion = window.matchMedia(
        '(prefers-reduced-motion: reduce)'
      ).matches;

      context = gsap.context(() => {
        const memberCards =
          gsap.utils.toArray<HTMLElement>(
            '[data-member-card]',
            membersGridElement
          );

        const memberCardInners =
          gsap.utils.toArray<HTMLElement>(
            '[data-member-card-inner]',
            membersGridElement
          );

        const heroTextElements = [
          firstTitleElement,
          secondTitleElement,
          descriptionElement
        ].filter(
          (element): element is HTMLElement =>
            element !== undefined
        );

        if (reduceMotion) {
          gsap.set(
            [
              glowElement,
              decorationElement,
              ...heroTextElements,
              membersHeadingElement,
              programHeadingElement,
              programContentElement,
              ...memberCards,
              ...memberCardInners
            ],
            {
              clearProps: 'all'
            }
          );

          return;
        }

        gsap.set(glowElement, {
          opacity: 0,
          scale: 0.65
        });

        gsap.set(decorationElement, {
          opacity: 0,
          x: 100,
          rotate: 25,
          scale: 0.8
        });

        gsap.set(firstTitleElement, {
          y: 70,
          opacity: 0,
          scale: 0.92,
          filter: 'blur(12px)'
        });

        if (secondTitleElement) {
          gsap.set(secondTitleElement, {
            y: 75,
            opacity: 0,
            scale: 0.92,
            filter: 'blur(12px)'
          });
        }

        gsap.set(descriptionElement, {
          y: 45,
          opacity: 0,
          filter: 'blur(7px)'
        });

        const heroTimeline = gsap.timeline({
          defaults: {
            ease: 'power3.out'
          },
          scrollTrigger: {
            trigger: sectionElement,
            start: 'top 88%',
            once: true
          }
        });

        heroTimeline
          .to(
            glowElement,
            {
              opacity: 1,
              scale: 1,
              duration: 1.5
            },
            0
          )
          .to(
            decorationElement,
            {
              opacity: 1,
              x: 0,
              rotate: 45,
              scale: 1,
              duration: 1.3
            },
            0
          )
          .to(
            firstTitleElement,
            {
              y: 0,
              opacity: 1,
              scale: 1,
              filter: 'blur(0px)',
              duration: 1
            },
            0.15
          );

        if (secondTitleElement) {
          heroTimeline.to(
            secondTitleElement,
            {
              y: 0,
              opacity: 1,
              scale: 1,
              filter: 'blur(0px)',
              duration: 1
            },
            0.28
          );
        }

        heroTimeline
          .to(
            descriptionElement,
            {
              y: 0,
              opacity: 1,
              filter: 'blur(0px)',
              duration: 0.85
            },
            0.55
          )
          .set(heroTextElements, {
            clearProps:
              'opacity,transform,filter'
          });

        gsap.set(membersHeadingElement, {
          y: 55,
          opacity: 0,
          scale: 0.94,
          filter: 'blur(8px)'
        });

        gsap.set(memberCards, {
          y: 60,
          opacity: 0,
          scale: 0.94
        });

        gsap.set(memberCardInners, {
          rotateY: 180,
          transformOrigin: 'center center',
          transformPerspective: 1200
        });

        const membersTimeline = gsap.timeline({
          defaults: {
            ease: 'power3.out'
          },
          scrollTrigger: {
            trigger: membersSectionElement,
            start: 'top 80%',
            once: true
          }
        });

        membersTimeline
          .to(membersHeadingElement, {
            y: 0,
            opacity: 1,
            scale: 1,
            filter: 'blur(0px)',
            duration: 0.8
          })
          .to(
            memberCards,
            {
              y: 0,
              opacity: 1,
              scale: 1,
              duration: 0.45,
              stagger: 0.16
            },
            0.25
          )
          .to(
            memberCardInners,
            {
              rotateY: 0,
              duration: 1.15,
              stagger: 0.16,
              ease: 'power3.inOut'
            },
            0.35
          )
          .set(
            [
              membersHeadingElement,
              ...memberCards,
              ...memberCardInners
            ],
            {
              clearProps:
                'opacity,transform,filter,transformOrigin,transformPerspective'
            }
          );

        gsap.set(programHeadingElement, {
          y: 55,
          opacity: 0,
          scale: 0.94,
          filter: 'blur(8px)'
        });

        gsap.set(programContentElement, {
          y: 70,
          opacity: 0,
          scale: 0.98
        });

        const programTimeline = gsap.timeline({
          defaults: {
            ease: 'power3.out'
          },
          scrollTrigger: {
            trigger: programSectionElement,
            start: 'top 80%',
            once: true
          }
        });

        programTimeline
          .to(programHeadingElement, {
            y: 0,
            opacity: 1,
            scale: 1,
            filter: 'blur(0px)',
            duration: 0.85
          })
          .to(
            programContentElement,
            {
              y: 0,
              opacity: 1,
              scale: 1,
              duration: 0.75
            },
            0.25
          )
          .set(
            [
              programHeadingElement,
              programContentElement
            ],
            {
              clearProps:
                'opacity,transform,filter'
            }
          );
      }, sectionElement);

      ScrollTrigger.refresh();

      document.documentElement.style.scrollBehavior =
        previousScrollBehavior;
    }

    forceScrollToTop();

    const firstFrame = requestAnimationFrame(() => {
      forceScrollToTop();

      secondFrame = requestAnimationFrame(() => {
        forceScrollToTop();
        setupAnimations();
      });
    });

    return () => {
      cancelAnimationFrame(firstFrame);
      cancelAnimationFrame(secondFrame);

      context?.revert();

      window.history.scrollRestoration =
        previousScrollRestoration;

      document.documentElement.style.scrollBehavior =
        previousScrollBehavior;
    };
  });
</script>

<svelte:head>
  <title>{title}</title>

  <meta
    name="description"
    content={description}
  />
</svelte:head>

<section
  bind:this={sectionElement}
  class="
    relative min-h-screen overflow-hidden
    bg-gradient-to-b
    from-blue-950
    via-blue-400
    to-blue-50
    pb-28 pt-14
    sm:pt-20
    lg:pt-28
    lg:pb-36 
  "
>
  <div
    bind:this={glowElement}
    class="
      pointer-events-none absolute
      left-1/2 top-8
      h-[500px] w-[900px]
      -translate-x-1/2
      rounded-full
      bg-blue-300/10
      blur-[130px]
    "
  ></div>

  <div
    bind:this={decorationElement}
    class="
      pointer-events-none absolute
      -right-32 top-24
      size-96 rotate-45
      border border-blue-200/10
    "
  ></div>

  <div
    class="
      relative z-10 mx-auto flex
      w-full max-w-7xl
      flex-col items-center
      text-center
    "
  >
    <div
      class="
        flex w-full max-w-5xl
        flex-col items-center
      "
    >
      <h1
        class="
          text-[clamp(2.8rem,8vw,7rem)]
          leading-[0.92] font-black
          tracking-[-0.065em]
        "
      >
        <span
          bind:this={firstTitleElement}
          class="
            block text-blue-600
            [-webkit-text-stroke:2px_white]
            [paint-order:stroke_fill]
            drop-shadow-[0_10px_0_rgba(30,64,175,0.55)]
            sm:[-webkit-text-stroke:3px_white]
          "
        >
          {titleLines.first}
        </span>

        {#if titleLines.second}
          <span
            bind:this={secondTitleElement}
            class="
              block text-blue-800
              [-webkit-text-stroke:2px_white]
              [paint-order:stroke_fill]
              drop-shadow-[0_10px_0_rgba(5,34,70,0.78)]
              sm:[-webkit-text-stroke:3px_white]
            "
          >
            {titleLines.second}
          </span>
        {/if}
      </h1>

      <p
        bind:this={descriptionElement}
        class="
          mx-auto mt-8 max-w-[70%]
          text-sm leading-6 font-medium
          text-white
          drop-shadow-[0_2px_2px_rgba(0,0,0,0.55)]
          sm:text-base sm:leading-7
          md:text-lg md:leading-8
        "
      >
        {#each descriptionParts as part, index (`${index}-${part.text}`)}
          <span
            class={part.highlighted
              ? 'font-black text-white'
              : ''}
          >
            {part.text}
          </span>
        {/each}
      </p>
    </div>

    <div
      bind:this={membersSectionElement}
      class="mt-20 w-full sm:mt-24"
    >
      <h2
        bind:this={membersHeadingElement}
        class="
          text-4xl leading-none font-black
          tracking-[-0.055em]
          text-blue-700
          [-webkit-text-stroke:1.5px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_8px_0_rgba(5,34,70,0.55)]
          sm:text-5xl
          lg:text-6xl
        "
      >
        Struktur Kementerian
      </h2>

      <div
        bind:this={membersGridElement}
        class="
          mx-auto mt-14 grid
          w-full max-w-6xl
          grid-cols-1 items-center
          gap-10
          md:grid-cols-3
          md:gap-8
          lg:gap-14
        "
      >
        {#each members as member (member.id)}
          <div
            class={member.featured
              ? 'order-first md:order-none'
              : ''}
          >
            <MemberCard
              role={member.role}
              title={member.title}
              image={member.image}
              featured={member.featured}
            />
          </div>
        {/each}
      </div>
    </div>

    <div
      bind:this={programSectionElement}
      class="mt-28 w-full sm:mt-36"
    >
      <h2
        bind:this={programHeadingElement}
        class="
          text-4xl leading-none font-black
          tracking-[-0.055em]
          text-blue-700
          [-webkit-text-stroke:1.5px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_8px_0_rgba(5,34,70,0.55)]
          sm:text-5xl
          lg:text-6xl
        "
      >
        Program Kerja
      </h2>

      <div
        bind:this={programContentElement}
        class="mt-12 w-full"
      >
        <ListProker
          {programs}
          pageSize={8}
          queryKey="prokerPage"
        />
      </div>
    </div>
  </div>

   <img src="/menko/b-3-left.png" alt="bintang 3 left" class="size-28 md:size-40 absolute top-1/3 left-0" />
  <img src="/menko/b-3-left.png" alt="bintang 3 left" class="size-28 md:size-40 absolute top-2/3 left-0" />
  <img src="/menko/b-3-right.png" alt="bintang 3 right" class="size-28 md:size-40 absolute top-1/5 right-0" />
</section>