### Enterprise & Organizing

> "Cloudy with High Chance of DBMS: A 10-year Prediction for Enterprise-Grade ML" -- http://cidrdb.org/cidr2020/papers/p8-agrawal-cidr20.pdf

> "For the longest time, we fixated on cycles, making things work faster—the processors, CPUs, GPUs, and more parallel servers. We ignored the data part. Now we have to fixate on data." -- https://www.technologyreview.com/2021/10/27/1037134/data-science-challenges-trustworthy-ai/

> "the teams that produce data are decoupled from the downstream analytics and ML use cases. Likewise, the people dealing with analytics use cases are typically unaware of operational systems that produce the data" -- https://m.subbu.org/broken-state-of-data-4c8a8a30a0c3

> "Towards Observability for Machine Learning Pipelines" -- https://arxiv.org/pdf/2108.13557.pdf

> [The Future Demands Every Company to be an AI Company](https://medium.com/@shanksphere/the-future-demands-every-company-to-be-an-ai-company-a352d3b83b0d)
> 
> Accessible and Explainable AI for Modern Data Teams - https://app.chaya.ai/
>  

Visualization of Data Science workflow https://docs.metaflow.org/introduction/why-metaflow

- Citizen Data Science - https://future.a16z.com/solo-workers-software-stack/
  - https://databricks.com/blog/2021/10/06/bringing-lakehouse-to-the-citizen-data-scientist-announcing-the-acquisition-of-8080-labs.html
  - https://www.nanalyze.com/2021/04/no-code-platforms-machine-learning/

- Open (Public?) Datasets Hub:
  - https://github.com/quiltdata/quilt 

### Foundation Models

> But the most important trend I want to comment on is that the whole setting of training a neural network from scratch on some target task (like digit recognition) is quickly becoming outdated due to finetuning, especially with the emergence of foundation models like GPT. These foundation models are trained by only a few institutions with substantial computing resources, and most applications are achieved via lightweight finetuning of part of the network, prompt engineering, or an optional step of data or model distillation into smaller, special-purpose inference networks. I think we should expect this trend to be very much alive, and indeed, intensify. In its most extreme extrapolation, you will not want to train any neural networks at all. In 2055, you will ask a 10,000,000X-sized neural net megabrain to perform some task by speaking (or thinking) to it in English. And if you ask nicely enough, it will oblige. Yes you could train a neural net too… but why would you? -- [Deep Neural Nets: 33 years ago and 33 years from now](http://karpathy.github.io/2022/03/14/lecun1989/) 

- https://www.thebulwark.com/how-ai-is-being-transformed-by-foundation-models/

--

* Database / Warehouse:
  * https://www.firebolt.io/resources/cloud-data-warehouse-comparison

* Wrangling:
  * https://www.getdbt.com/

* Catalogs:
  * [Awesome Data Discovery and Observability](https://github.com/opendatadiscovery/awesome-data-catalogs)
    * [Discover, Collaborate and get your Data Right](https://open-metadata.org/) / Select Star, DataHub (Acryl Data), Mataphor Data, Stemma, Alvin, ...
  * [Smart Data Models](https://smartdatamodels.org/)
  
* Private data:
> Organizations have vast amounts of confidential data locked down due to privacy concerns. Opaque Systems makes confidential data useful by enabling secure analytics and machine learning on encrypted data that comes from one or more sources.
With Opaque Systems, organizations can analyze encrypted data in the cloud using popular tools like Apache Spark, while ensuring that their data is never exposed unencrypted to the cloud provider. -- https://opaque.co/

> Confidential computing is most commonly used in financial services and health care industries and by government agencies, but every industry can benefit from it. -- https://azure.microsoft.com/en-us/solutions/confidential-compute/#customer-stories

--

OpenAI:
* [Applications](https://beta.openai.com/examples)
* [Can Code](https://spectrum.ieee.org/openai-wont-replace-coders)
> "due to the cost of training, it wasn't feasible to retrain the model."
[Cost of Training](https://spectrum.ieee.org/deep-learning-computational-cost)
* [AI and Compute](https://openai.com/blog/ai-and-compute/)
* [Cohere](https://venturebeat.com/2021/11/15/openai-rival-cohere-launches-language-model-api/)

Stanford:
  * https://sisudata.com/
  * https://ai.stanford.edu/blog/
  * https://hai.stanford.edu/
  * https://crfm.stanford.edu/
> In recent years, a new successful paradigm for building AI systems has emerged: Train one model on a huge amount of data and adapt it to many applications. We call such a model a foundation model https://www.youtube.com/watch?v=RLrjKGN89Fc
  * https://stanford-cs329s.github.io/syllabus.html

University of Washington:
  * https://falx.cs.washington.edu/

UC Berkeley:
  * [After Databricks](https://rise.cs.berkeley.edu/blog/the-inside-story-of-how-uc-berkeley-became-the-incubator-for-red-hot-enterprise-startups-databricks-sifive-and-anyscale/)
    * [Anyscale](https://docs.ray.io/en/master/index.html)

CMU:
  * https://ottertune.com/
  * https://db.cs.cmu.edu/seminar2021-dose2/

https://incidentdatabase.ai/

---

Modern Data Stack - https://sisudata.com/blog/modern-analytics-stack / https://www.moderndatastack.xyz/companies

* https://rudderstack.com/blog/the-state-of-data-engineering-in-2022/
* MAD (Machine Learning, AI & Data) Crowded Landscape, 2021: https://mattturck.com/2021top10/ 
* Linux Foundation Landscape https://landscape.lfai.foundation/ 

* Databases, 2021: https://ottertune.com/blog/2021-databases-retrospective/

* Dominance, Companies being built on top of As a Platform (Cloud - AWS, Azure & GCP)
  * Snowflake
  * Databricks

* Growth of OSS:

  * https://imply.io/
  * https://vectorized.io/
  * https://www.decodable.co/
  * https://materialize.com/

  * https://airbyte.com/
  * https://www.getdbt.com/
  * https://dagster.io/
  * https://www.astronomer.io/
  * https://www.prefect.io/
  * https://chronosphere.io/

  * https://greatexpectations.io/
  * https://www.soda.io/

  * https://databricks.com/product/unity-catalog
  * https://datahubproject.io/ 
  * https://open-metadata.org/

  * https://sisudata.com/
  * https://www.anyscale.com/

  * https://observablehq.com/

* Notebooks: Jupyter was awarded the 2017 ACM Software Systems Award — a prestigious honor it shares with Java, Unix, and the Web.
  * [SQL Notebooks: Combining the power of Jupyter and SQL editors for data analytics](https://engineering.fb.com/2022/04/26/developer-tools/sql-notebooks/)
  * [JupyterLite is the Jupyter notebook stack using Python-compiled-to-WebAssembly, running directly in your browser, no server required](https://jupyter.org/try)
  * https://netflixtechblog.com/notebook-innovation-591ee3221233
  * https://code.visualstudio.com/blogs/2021/08/05/notebooks / https://code.visualstudio.com/blogs/2021/11/08/custom-notebooks
  * https://datalore.jetbrains.com/
  * https://docs.databricks.com/notebooks/index.html
  * Vertex AI - https://cloud.google.com/blog/topics/retail/retail-forecasting-ai-real-time-insights-vertex-ai-forecast
  * Sagemaker - https://studiolab.sagemaker.aws/
  * Azure Machine Learning Studio - 
