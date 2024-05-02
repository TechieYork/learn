import { Col, Layout, Row, Flex, Typography } from "antd";
import { useNavigate } from "react-router-dom";
import { CourseCard } from "../../components/CourseCard";
import { Intro } from "../../components/Introduction";
import { UOWCAI_COURSES } from "../../constants/Course";
const { Title } = Typography;

function Home() {
  let navigate = useNavigate();

  return (
    <Layout.Content>
      <Row justify="center">
        <Col span={2} />
        <Col span={20}>
          <Intro />
        </Col>
        <Col span={2} />
      </Row>
      <Row justify="center">
        <Col span={2} />
        <Col span={20}>
          <Title level={3}>Latest Courses</Title>
          <Row gutter={[16, 16]}>
            {UOWCAI_COURSES.map((course, index) => (
              <Flex
                wrap="wrap"
                gap="middle"
                flex="1 0 30%"
                style={{ margin: "10px" }}
              >
                <CourseCard
                  title={course.title}
                  description={course.description}
                  image={course.image}
                  url={course.url}
                  onClick={() => navigate(course.url)}
                />
              </Flex>
            ))}
          </Row>
        </Col>
        <Col span={2} />
      </Row>
    </Layout.Content>
  );
}

export { Home };
