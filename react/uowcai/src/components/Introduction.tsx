import { Divider, Typography } from "antd";
import {
  UOWCAI,
  UOWCAI_DESCRIPTION,
  UOWCAI_MISSION,
} from "../constants/Course";

const { Title, Paragraph, Text, Link } = Typography;

function Intro() {
  return (
    <Typography>
      <Title level={2}>{UOWCAI}</Title>
      <Paragraph>{UOWCAI_DESCRIPTION}</Paragraph>
      <Title level={3}>Our Mission</Title>
      {UOWCAI_MISSION}
      <Title level={3}>Who will benefit</Title>
      <Text>
        The series is specifically designed for the following cohort of
        professionals{" "}
      </Text>{" "}
      <Text italic strong>
        {" "}
        who are not from Computer Science and Data Science disciplines.
      </Text>
      <ul style={{textIndent: "10px"}}>
        <li>
          Staff (academic and professional) and High Degree Research students
          who are interested in applying AI, ML and/or DL in their research and
          services.
        </li>
        <li>
          Professionals from industries, government departments and communities
          who are interested in understanding and applying AI, ML and/or DL in
          their products and services.
        </li>
      </ul>
    </Typography>
  );
}

export { Intro };
