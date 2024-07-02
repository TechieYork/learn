// UOW
const UOW_LOGO = process.env.PUBLIC_URL + "/assets/uow-logo.png";

// Introduction of Home Page
const UOWCAI = "Centre for Artificial Intelligence (CAI)";
const UOWCAI_DESCRIPTION = `Welcome to the official GitHub repository for the Centre for Artificial Intelligence (CAI) at the University of Wollongong. CAI conducts pioneering research to develop innovative theories and techniques of AI, and transferring the knowledge and technologies to industry, community and society.`;
const UOWCAI_MISSION = `To conduct pioneering research to develop innovative theories and techniques of AI, and transferring the knowledge and technologies to industry, community and society.`;

const UOWCAI_COURSE_DEFAULT_IMAGE =
  process.env.PUBLIC_URL + "/assets/course-default.png";

// Course Information
const UOWCAI_COURSES = [
  {
    id: "gai",
    title: "Generative AI",
    description: `Machine learning is a method of data analysis that automates analytical model building. It is a branch of artificial intelligence based on the idea that systems can learn from data, identify patterns and make decisions with minimal human intervention.`,
    image: process.env.PUBLIC_URL + "/assets/ml.png",
    url: "/courses/gai",
    learning: "/learning/gai",
    outcomes: [
      "Understand the basic concepts of machine learning",
      "Understand the mathematical foundations of machine learning",
      "Understand the practical applications of machine learning",
    ],
    pdf:
      process.env.PUBLIC_URL +
      "/pdfs/A Gentle Introduction to Generative AI.pdf",
  },
  {
    id: "ml",
    title: "Machine Learning",
    description: `Machine learning is a method of data analysis that automates analytical model building. It is a branch of artificial intelligence based on the idea that systems can learn from data, identify patterns and make decisions with minimal human intervention.`,
    image: process.env.PUBLIC_URL + "/assets/ml.png",
    url: "/courses/ml",
    learning: "/learning/ml",
    outcomes: [
      "Understand the basic concepts of machine learning",
      "Understand the mathematical foundations of machine learning",
      "Understand the practical applications of machine learning",
    ],
    pdf: process.env.PUBLIC_URL + "/pdfs/A Gentle Introduction to Contemporary AI.pdf",
  },
  {
    id: "cv",
    title: "Computer Vision",
    description: `Computer vision is an interdisciplinary scientific field that deals with how computers can gain high-level understanding from digital images or videos. From the perspective of engineering, it seeks to understand and automate tasks that the human visual system can do.`,
    image: process.env.PUBLIC_URL + "/assets/cv.jpeg",
    url: "/courses/cv",
    learning: "/learning/cv",
    outcomes: [
      "Understand the basic concepts of computer vision",
      "Understand the mathematical foundations of computer vision",
      "Understand the practical applications of computer vision",
    ],
  },
  {
    id: "nlp",
    title: "Natural Language Processing",
    description: `Natural language processing (NLP) is a subfield of linguistics, computer science, and artificial intelligence concerned with the interactions between computers and human language, in particular how to program computers to process and analyze large amounts of natural language data.`,
    image: process.env.PUBLIC_URL + "/assets/nlp.png",
    url: "/courses/nlp",
    learning: "/learning/nlp",
    outcomes: [
      "Understand the basic concepts of natural language processing",
      "Understand the mathematical foundations of natural language processing",
      "Understand the practical applications of natural language processing",
    ],
  },
];

const UOWCAI_COURSES_OUTLINE = {
  gai: {
    id: "gai",
    outline: [
      {
        name: "Introduction",
        start: 1,
        topics: [
          {
            name: "Centre for Artificial Intelligence (CAI)",
            page: 2,
          },
          {
            name: "Brief Outline",
            page: 3,
          },
        ],
      },
      {
        name: "Generative Artificial Intelligence (GAI)",
        start: 4,
        topics: [
          {
            name: "Definition of GAI",
            page: 4,
          },
          {
            name: "Predictive/Descriptive AI vs Generative AI (GAI)",
            page: 5,
          },
          {
            name: "How GAI Works? Training, prompt and inference (generation)",
            page: 6,
          },
          {
            name: "How GAI Works? From prompt to output",
            page: 7,
          },
        ],
      },
    ],
  },
  ml: {
    id: "ml",
    outline: [
      {
        name: "Introduction",
        start: 1,
        topics: [
          {
            name: "Centre for Artificial Intelligence (CAI)",
            page: 2,
          },
          {
            name: "Brief Outline",
            page: 3,
          },
        ],
      },
      {
        name: "Machine Learning",
        start: 4,
        topics: [
          {
            name: "Definition of GAI",
            page: 4,
          },
          {
            name: "Predictive/Descriptive AI vs Generative AI (GAI)",
            page: 5,
          },
          {
            name: "How GAI Works? Training, prompt and inference (generation)",
            page: 6,
          },
          {
            name: "How GAI Works? From prompt to output",
            page: 7,
          },
        ],
      },
    ],
  },
};

export {
  UOW_LOGO,
  UOWCAI,
  UOWCAI_DESCRIPTION,
  UOWCAI_MISSION,
  UOWCAI_COURSES,
  UOWCAI_COURSE_DEFAULT_IMAGE,
  UOWCAI_COURSES_OUTLINE,
};
