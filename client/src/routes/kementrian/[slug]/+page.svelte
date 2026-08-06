<script lang="ts">
  import { onMount,  } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  import MemberCard from '$lib/components/kementrian/MemberCard.svelte';
  import ListProker from '$lib/components/kementrian/ListProker.svelte';

  import type {
    OrganizationMember,
    WorkProgram
  } from '$lib/types';
  import type { PageData } from './$types';

  let {
    data
  }: {
    data: PageData;
  } = $props();

  $effect(() => console.log(data));

  type DescriptionPart = {
    text: string;
    highlighted: boolean;
  };

  type Member = {
    id: number;
    role: string;
    title: string;
    image: string;
    featured: boolean;
  };

  type ProgramKerja = {
    id: string;
    title: string;
    description: string;
    slug: string;
  };

  const unit = $derived(data.unit);

  const title = $derived(
    unit.name ??
      unit.short_name ??
      'Kementerian'
  );

  const description = $derived(
    unit.description ??
      `${title} merupakan bagian dari Kabinet Cerita Loka BEM Universitas Airlangga 2026.`
  );

  function getMemberImage(
    member: OrganizationMember
  ): string {
    return (
      member.photo?.url ??
      member.photo?.thumbnail_url ??
      '/kementrian/test.png'
    );
  }

  function getMemberRole(
    member: OrganizationMember
  ): string {
    const roles: Record<string, string> = {
      MENKO: 'Menteri Koordinator',
      MENTERI: 'Menteri',
      MINISTER: 'Menteri',
      DIRJEN: 'Direktur Jenderal',
      DIRECTOR_GENERAL: 'Direktur Jenderal',
      DEPUTI: 'Deputi',
      KEPALA: 'Kepala',
      KABIRO: 'Kepala Biro'
    };

    return (
      roles[member.position_type] ??
      member.position_type
    );
  }

  const members = $derived.by<Member[]>(() => {
    return [...(unit.members ?? [])]
      .filter(
        (member) =>
          member.is_active !== false
      )
      .sort((first, second) => {
        if (
          first.is_leader !==
          second.is_leader
        ) {
          return first.is_leader ? -1 : 1;
        }

        return (
          (first.display_order ?? 0) -
          (second.display_order ?? 0)
        );
      })
      .map((member) => ({
        id: member.id,
        role: getMemberRole(member),
        title: member.name,
        image: getMemberImage(member),
        featured: member.is_leader
      }));
  });

  function getProgramDescription(
    program: WorkProgram
  ): string {
    return (
      program.short_description ??
      program.description ??
      'Informasi program kerja akan segera diperbarui.'
    );
  }

  const programs =
    $derived.by<ProgramKerja[]>(() => {
      return [...(data.programs ?? [])]
        .filter(
          (program) =>
            program.is_published !== false
        )
        .sort(
          (first, second) =>
            (first.display_order ?? 0) -
            (second.display_order ?? 0)
        )
        .map((program) => ({
          id: String(program.id),
          title: program.name,
          description:
            getProgramDescription(program),
          slug: program.slug
        }));
    });

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
        description.toLocaleLowerCase(
          'id-ID'
        );

      const lowerTitle =
        normalizedTitle.toLocaleLowerCase(
          'id-ID'
        );

      const parts: DescriptionPart[] = [];
      let currentIndex = 0;

      while (
        currentIndex < description.length
      ) {
        const matchIndex =
          lowerDescription.indexOf(
            lowerTitle,
            currentIndex
          );

        if (matchIndex === -1) {
          parts.push({
            text: description.slice(
              currentIndex
            ),
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
            matchIndex +
              normalizedTitle.length
          ),
          highlighted: true
        });

        currentIndex =
          matchIndex +
          normalizedTitle.length;
      }

      return parts;
    });

  let sectionElement =
    $state<HTMLElement>();

  let glowElement =
    $state<HTMLDivElement>();

  let decorationElement =
    $state<HTMLDivElement>();

  let firstTitleElement =
    $state<HTMLSpanElement>();

  let secondTitleElement =
    $state<HTMLSpanElement>();

  let descriptionElement =
    $state<HTMLParagraphElement>();

  let membersSectionElement =
    $state<HTMLDivElement>();

  let membersHeadingElement =
    $state<HTMLHeadingElement>();

  let membersGridElement =
    $state<HTMLDivElement>();

  let programSectionElement =
    $state<HTMLDivElement>();

  let programHeadingElement =
    $state<HTMLHeadingElement>();

  let programContentElement =
    $state<HTMLDivElement>();

  onMount(() => {
    gsap.registerPlugin(ScrollTrigger);

    const previousScrollRestoration =
      window.history.scrollRestoration;

    const previousScrollBehavior =
      document.documentElement.style
        .scrollBehavior;

    let context:
      | gsap.Context
      | undefined;

    let secondFrame = 0;

    window.history.scrollRestoration =
      'manual';

    document.documentElement.style
      .scrollBehavior = 'auto';

    function forceScrollToTop(): void {
      window.scrollTo({
        top: 0,
        left: 0,
        behavior: 'auto'
      });

      document.documentElement.scrollTop =
        0;

      document.body.scrollTop = 0;
    }

    function setupAnimations(): void {
      const section = sectionElement;
      const glow = glowElement;
      const decoration = decorationElement;
      const firstTitle = firstTitleElement;
      const secondTitle =
        secondTitleElement;
      const descriptionText =
        descriptionElement;
      const membersSection =
        membersSectionElement;
      const membersHeading =
        membersHeadingElement;
      const membersGrid =
        membersGridElement;
      const programSection =
        programSectionElement;
      const programHeading =
        programHeadingElement;
      const programContent =
        programContentElement;

      if (
        !section ||
        !glow ||
        !decoration ||
        !firstTitle ||
        !descriptionText ||
        !membersSection ||
        !membersHeading ||
        !membersGrid ||
        !programSection ||
        !programHeading ||
        !programContent
      ) {
        return;
      }

      const reduceMotion =
        window.matchMedia(
          '(prefers-reduced-motion: reduce)'
        ).matches;

      context = gsap.context(() => {
        const memberCards =
          gsap.utils.toArray<HTMLElement>(
            '[data-member-card]',
            membersGrid
          );

        const memberCardInners =
          gsap.utils.toArray<HTMLElement>(
            '[data-member-card-inner]',
            membersGrid
          );

        const heroTextElements = [
          firstTitle,
          secondTitle,
          descriptionText
        ].filter(
          (
            element
          ): element is HTMLElement =>
            element !== undefined
        );

        if (reduceMotion) {
          gsap.set(
            [
              glow,
              decoration,
              ...heroTextElements,
              membersHeading,
              programHeading,
              programContent,
              ...memberCards,
              ...memberCardInners
            ],
            {
              clearProps: 'all'
            }
          );

          return;
        }

        gsap.set(glow, {
          opacity: 0,
          scale: 0.65
        });

        gsap.set(decoration, {
          opacity: 0,
          x: 100,
          rotate: 25,
          scale: 0.8
        });

        gsap.set(firstTitle, {
          y: 70,
          opacity: 0,
          scale: 0.92,
          filter: 'blur(12px)'
        });

        if (secondTitle) {
          gsap.set(secondTitle, {
            y: 75,
            opacity: 0,
            scale: 0.92,
            filter: 'blur(12px)'
          });
        }

        gsap.set(descriptionText, {
          y: 45,
          opacity: 0,
          filter: 'blur(7px)'
        });

        const heroTimeline =
          gsap.timeline({
            defaults: {
              ease: 'power3.out'
            },
            scrollTrigger: {
              trigger: section,
              start: 'top 88%',
              once: true
            }
          });

        heroTimeline
          .to(
            glow,
            {
              opacity: 1,
              scale: 1,
              duration: 1.5
            },
            0
          )
          .to(
            decoration,
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
            firstTitle,
            {
              y: 0,
              opacity: 1,
              scale: 1,
              filter: 'blur(0px)',
              duration: 1
            },
            0.15
          );

        if (secondTitle) {
          heroTimeline.to(
            secondTitle,
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
            descriptionText,
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

        gsap.set(membersHeading, {
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
          transformOrigin:
            'center center',
          transformPerspective: 1200
        });

        const membersTimeline =
          gsap.timeline({
            defaults: {
              ease: 'power3.out'
            },
            scrollTrigger: {
              trigger: membersSection,
              start: 'top 80%',
              once: true
            }
          });

        membersTimeline
          .to(membersHeading, {
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
              membersHeading,
              ...memberCards,
              ...memberCardInners
            ],
            {
              clearProps:
                'opacity,transform,filter,transformOrigin,transformPerspective'
            }
          );

        gsap.set(programHeading, {
          y: 55,
          opacity: 0,
          scale: 0.94,
          filter: 'blur(8px)'
        });

        gsap.set(programContent, {
          y: 70,
          opacity: 0,
          scale: 0.98
        });

        const programTimeline =
          gsap.timeline({
            defaults: {
              ease: 'power3.out'
            },
            scrollTrigger: {
              trigger: programSection,
              start: 'top 80%',
              once: true
            }
          });

        programTimeline
          .to(programHeading, {
            y: 0,
            opacity: 1,
            scale: 1,
            filter: 'blur(0px)',
            duration: 0.85
          })
          .to(
            programContent,
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
              programHeading,
              programContent
            ],
            {
              clearProps:
                'opacity,transform,filter'
            }
          );
      }, section);

      ScrollTrigger.refresh();

      document.documentElement.style
        .scrollBehavior =
        previousScrollBehavior;
    }

    forceScrollToTop();

    const firstFrame =
      requestAnimationFrame(() => {
        forceScrollToTop();

        secondFrame =
          requestAnimationFrame(() => {
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

      document.documentElement.style
        .scrollBehavior =
        previousScrollBehavior;
    };
  });
</script>

<svelte:head>
  <title>{title} | BEM UNAIR</title>

  <meta
    name="description"
    content={description}
  />
</svelte:head>

<section
  bind:this={sectionElement}
  class="
    relative min-h-screen overflow-hidden
    bg-linear-to-b
    from-blue-950
    via-blue-400
    to-blue-50
    px-5 pb-28 pt-14
    sm:px-8 sm:pt-20
    lg:px-12 lg:pb-36 lg:pt-28
  "
>
  <div
    bind:this={glowElement}
    class="
      pointer-events-none absolute
      left-1/2 top-8
      h-125 w-225
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
          mx-auto mt-8 max-w-4xl
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
        {#if members.length > 0}
          {#each members as member (member.id)}
            <div
              class={member.featured
                ? 'md:col-start-2 md:row-start-1'
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
        {:else}
          <p
            class="
              col-span-full py-12
              text-base font-semibold
              text-white
            "
          >
            Data struktur kementerian belum tersedia.
          </p>
        {/if}
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
        {#if programs.length > 0}
          <ListProker
            {programs}
            pageSize={8}
            queryKey="prokerPage"
          />
        {:else}
          <p
            class="
              py-16 text-center
              text-base font-semibold
              text-blue-950/70
            "
          >
            Program kerja belum tersedia.
          </p>
        {/if}
      </div>
    </div>
  </div>

  <img
    src="/menko/b-3-left.png"
    alt=""
    class="
      pointer-events-none absolute
      left-0 top-1/3
      size-28
      md:size-40
    "
  />

  <img
    src="/menko/b-3-left.png"
    alt=""
    class="
      pointer-events-none absolute
      left-0 top-2/3
      size-28
      md:size-40
    "
  />

  <img
    src="/menko/b-3-right.png"
    alt=""
    class="
      pointer-events-none absolute
      right-0 top-1/5
      size-28
      md:size-40
    "
  />
</section>