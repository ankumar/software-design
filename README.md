1. Applications ( Modernization ) & Databases 
2. Analytics, AI & Machine Learning
3. Security, Cloud Platforms & Infrastructure
4. Retail, Manufacturing, Finance, Entertainment, Healthcare, ...

> There's a saying that writing software is more like tending a garden than constructing a building -- things constantly change.
>
> But the more I learn about how buildings evolve, I think this process is actually a perfect analogy for designing software!
> 
> Thread: https://twitter.com/geoffreylitt/status/1272542423001022467 
>
[![Overnight Success… 10 years](https://github.com/ankumar/architecture/blob/main/images/VS%20Code%20an%20Overnight%20Success.png)](https://www.youtube.com/watch?v=hilznKQij7A&list=PLj6YeMhvp2S6uB23beQaffszlavLq3lNq) 
[Software keeps Evolving and is Never Done](https://www.youtube.com/watch?v=4-0fiuKdxmc), first Devtools evolving like [VS Code 10 years journey By Erich Gamma, The "Gang of Four"](https://www.youtube.com/watch?v=hilznKQij7A&list=PLj6YeMhvp2S6uB23beQaffszlavLq3lNq). More design evolution in-progress at JetBrains with [Space](https://www.jetbrains.com/remote-development/space-dev-environments/) / [Gateway](https://www.jetbrains.com/remote-development/gateway/) / [Fleet](https://www.jetbrains.com/fleet/) & [Toolchain](https://github.com/ankumar/Open-software-design/blob/main/Radar/Devtools.md). Evolution are always fascinating from learning perspective, like all over night successes, 10+ years and 100s of man years of work! 

Switching to building software using Devtools, design a sustainable model with core principle of **Going far, Together**

- Design/UX, Architecture, DevOps/SRE, Quality, ...

> Having a community that you can learn from, … I think it was the reason I got into software engineering, the reason I was able to get into it so quickly and so easily. There’s so many different ways you can learn. -Kate (chemical) -- [The Crossover Project](https://www.hillelwayne.com/post/what-we-can-learn/)

- Inner Source, Open Source, ...

> think a lot of it we benefit a ton from the open source community and just all the learnings there that are laid bare in the open, all the mistakes, all the success, all the problems, it's a very slow moving process, usually open source, but it's very deliberate. And you get to see because of the the pace, you get to see what it takes to really build something meaningful, learned most most of everything I learned about hacking and programming and engineering has been due to open source and the the generosity that people have given to give up their time, sacrifice their time without any expectation in return other than being a part of something much larger than themselves. Yeah, I think it's great. -- [Engineering at scale](https://www.youtube.com/watch?v=60KJz1BVTyU&t=168s)  

**1. Seeding / Learning / Changing:**

> Videos: 
> 1. [Programming as architecture, design and urban planning](https://vimeo.com/669587851) / [PDF](http://tomasp.net/academic/papers/metaphors/metaphors.pdf)
> 2. [How Buildings Learn](https://www.youtube.com/user/brandst/videos)

- ["Ad Hoc is an experiment in combining outside experts and cutting edge technology with the amazing civil servants in government to create services that are not just functional but exceptional."](https://adhoc.team/about/)  

> Open-source software has made it easier for software developers to study and learn programming by looking at real-world working software. But what about software design? Building architects (not software) study thousands of buildings in their training and career. Software developers, in contrast, only study a handful of other large software designs. This means we repeat mistakes that we would have been able to avoid, should we be able to learn from others.
>
> Open software design is the idea that software design must be available for public use. It's the idea that **we should Work with the garage door up**. The main audience of the design is developers who built the software, and not external people. The contents are living artefacts, they naturally evolve as the software evolve. In the world where the majority of large real-world programs are closed-source, opening up our software design documentation is the least we can do to contribute back to the community.
>
> [Wisen's notes](https://notes.ceilfors.com/%C2%A7What's_top_of_mind.html)

- [Software for all: how do open-source communities work?](https://www.uoc.edu/portal/en/news/actualitat/2022/014-opens-source-software.html) 
- [“Crossovers”: people who used to work as traditional engineers and now work as software developers](https://www.hillelwayne.com/tags/crossover-project/)

- [AWS re:Invent 2021 - The architect elevator: Connecting IT and the boardroom](https://www.youtube.com/watch?v=nNbnXTl2VFQ) 

> once you empower people, giving them an environment in which to succeed, and recognise their successes, they will rapidly, and as a collective, start thinking about things which haven’t even crossed your mind. That’s the real benefit of this kind of approach: access to the collective intelligence of the many, over reliance on the much more restricted intelligence of the few.
>
> [Scaling the Practice of Architecture, Conversationally](https://martinfowler.com/articles/scaling-architecture-conversationally.html)

**2. Distributed:**

> [From 2021 Organizational Dynamics Masterclass](https://www.ruthmalan.com/) 
>
> [Polarity Management - Ex: Tactical/Strategic, Centralization/Decentralization](https://www.youtube.com/watch?v=yyuFr4gTzjU) 
>
> > Ruth Malan: Here's my understanding: In a polarity map, we recognize that there are poles, and over-emphasis could get us putting too much effort in one or other pole -- at which point the impetus to move away from the downsides of the pole would be strong. So we're seeking to get more "both and" which takes taking actions, and watching for the signals of a dip into too much of a pole (and hence too much of its downsides).  The polarities are unavoidable -- for example, we can't avoid there being some architecture decisions made in the teams without coordinating across teams. We want high autonomy in teams, so they're making a lot of decisions.  And some of these will turn out to be architectural(ly significant). At the same time as wanting teams to have degrees of independence, we want architecturally significant decisions to be made from a system perspective. So what can we do, to make more of that happen, and happen in time? (Whereas a dichotomy would say: pick one, you can't have both. If you have team autonomy you can't have architects (role or hat); if you have more team autonomy, you have less architecture (at least, less architecture (as intentional design; there will still be "accidental" architecture, etc.); etc. There is no notion that by taking actions, you could get more autonomy (where it matters) while still having (just enough) architecture (where it matters). Etc.) 
> >
> > Dana Bredemeyer: There are some interesting issues to tease apart here. Do you have a couple hours? :slightly_smiling_face: I find it helpful to distinguish between the level of an architecture and the level (granularity and location) of a decision. The level of the architecture can also be called it's scope or system view. Is the scope of a particular architecture the car, or the drive train (part of car), or the engine (part of drive train)? Let's say we're architecting the drive train. Drive train architectural decisions that do not impact car level desired outcomes are owned / made by the drive train team. If they DO affect car-level desired outcomes, then those decisions are owned / made by the car architecture team. Generally these are different teams, and the architectures are best viewed as different architectures. There's a car architecture, a drive train architecture, and an engine architecture. The DESIRED OUTCOMES cascade from broader to narrower scopes and back again. If the system being designed is a coherent system (such as a car) and all teams understand and are committed to overall car-level desired outcomes, then decisions at narrower scope must not compromise broader scope desired outcomes, but the need to achieve broader scope outcomes CAN drive very detailed decisions in narrower scope. All this gets more complicated when, for instance, the engine is made by a third party who also sells their engine for use in other cars, which have different sets of car-level desired outcome. Sheesh. We're just getting going here and already we can clearly see the highly-connected nature of organization strategy, product strategy, and organizational design. Architecture is challenging. I think Ackoff's short video also does a good job of addressing this. https://www.youtube.com/watch?v=OqEeIG8aPPk

- Developer Relations:
  - More from Shawn Wang, [Developer Relations](https://www.swyx.io/measuring-devrel/) 
  - Videos: https://www.youtube.com/c/Honeypotio/videos
   
[![Imagine this for an Enterprise's Cloud Infra:](https://github.com/ankumar/architecture/blob/main/images/Open%20Infrastructure%20Map.png)](https://openinframap.org/#8.78/37.7325/-121.3816) 
  
 --

1. Garage door Up:

[![Backstage Community](https://github.com/ankumar/architecture/blob/main/images/Acuity%20Brands%20Dev%20Portal.png)](https://backstage.io/demos)

Shared Context: Business, Product & Tech
- [Home Page](https://backstage.io/blog/2022/01/25/backstage-homepage-templates) -> **Explore:** Decision Records, Radar, Registry, ... **Create Components:** APIs, Docs, Tests, ... **Manage:** Deployments, Insights, ...

- Products
  - ["Agile" - Commerce Platform (Behind SSO)](https://agile.acuitybrandslighting.net/) -> Projects, Quotes, Orders, Claims, Returns, Distributor, Analytics, 
  - [Atrius - Buildings, Inside is where most people are most of the time](https://atrius.com/) 
 
- Design System 

- Code
  - Style Guides  
  - Tech Docs
  - Tests

 - Decision Records
  - Tech Radar
  - C4 Model

- Developer Experience
  - SaaS [Build/Buy](https://github.com/ankumar/Open-software-design/blob/main/Radar/SaaS.md#saas-alphabet) 
    - [Zeroheight](https://zeroheight.com/) [Bit](https://bit.dev/) [Storybook](https://storybook.js.org/) [ReadMe](https://readme.com/) 
    - [Figma](https://www.figma.com/) [Adobe Illustrator](https://creativecloud.adobe.com/) [Adobe Photoshop](https://creativecloud.adobe.com/) [Adobe Lightroom](https://creativecloud.adobe.com/) [QGIS](https://www.qgis.org/en/site/) [AutoCAD](https://web.autocad.com/login)
    - [Miro](https://miro.com/)
    - [Postman](https://www.postman.com/)
    - [Striim](https://www.striim.com/)
    - [Aqua](https://www.aquasec.com/)
    - [Datadog](https://www.datadoghq.com/)
    - [Google Cloud - BigQuery](https://cloud.google.com/)
    - [Azure - IAM, Compute & Databases](https://azure.microsoft.com/)
   
  - Open Source
    - [Cloud Relay API](https://github.com/DistechControls/CloudRelay)   
    - [Backstage](https://github.com/backstage/backstage) / SaaS: https://roadie.io/
    - [Temporal](https://github.com/temporalio/temporal)
    - [Artillery](https://github.com/artilleryio), [Playwright](https://github.com/microsoft/playwright)
    - [Argo CD](https://github.com/argoproj/argo-cd/)
  
2. Collaboration Opportunities:

> [12/16 10:27 AM] Quinn, Henry
>
> Good (but long) writeup on "why Temporal" from their Head of Developer Experience, Shawn Wang (aka Swyx).
> I've been following this guy for a while as he's been very vocal about learning in public and helped push a lot of folks (including myself) to tackle fun projects to learn new technologies while teaching others. I was very excited to see he joined Temporal! https://www.swyx.io/why-temporal/
>
>> [12/16 11:40 AM] Voils, Steven M
>> 
>> swyx has, presumably by happenstance, happened to show up on a number of podcasts i've been monitoring.  i agree that he's got a very refreshing take on how to learn, how to grow, and he's got a great story backing up his philosophy given his transition from finance to software. Quinn, Henry - this is actually a really great article (i'm perusing it) that i hadn't run into before.  i think it'd be a great thing to share on the blast i did in general, just to whet people's appetites for what's motivating us to do this kind of exploration in the first place

:arrow_up: A Thread inside a company **experimenting, learning & changing [1]**, an instance of organic influence from public domain & open source. The **cognitive load distributed [2]** & building context specific startegies is how we should think about acceleration of 10+ years and 100s of man years of work involved driving changes across organizations.

- Apps/State, Serverless Model & Sustainability - https://github.com/tufan-io/noun_and_verb / A first demo App: https://github.com/acuity-sr/nv-shopping-cart 

