import React from "react";
import {
  Col,
  Row,
  Image,
  Typography,
  Layout,
  Divider,
  Button,
  Card,
  Flex,
} from "antd";
import { useSelector, useDispatch } from "react-redux";
import { useParams, useNavigate } from "react-router-dom";
import { setCourseDetailImageLoaded } from "./courseDetailSlice";
import {
  UOWCAI_COURSES,
  UOWCAI_COURSE_DEFAULT_IMAGE,
} from "../../constants/Course";

const { Title, Paragraph, Text, Link } = Typography;

const TABS = [
  {
    label: "Introduction",
    key: "1",
  },
  {
    label: "Outcome",
    key: "2",
  },
  {
    label: "Comments",
    key: "3",
  },
];

interface CourseDetailProps {
  title?: "";
  image?: "";
  description?: "";
}

const CourseDetail = (props: CourseDetailProps) => {
  // get courseID from URL
  const courseDetailImageLoaded = useSelector(
    (state: any) => state.courseDetail.courseDetailImageLoaded
  );
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const params = useParams();
  const courseID = params.courseId;

  // get course info from storage
  let course = UOWCAI_COURSES.find((course) =>
    course.id === courseID ? course : null
  );
  if (course === undefined) {
    navigate("/not-found");
  }

  return (
    <Layout.Content style={{ zIndex: 100, margin: "3px" }}>
      <Row justify={"center"} style={{marginTop: "48px"}}>
        <Col span={2} />
        <Col span={20}>
          <Row justify={"center"}>
            <Flex wrap="wrap">
              <Flex
                wrap="wrap"
                gap="middle"
                flex="1 0 30%"
                style={{ margin: "10px" }}
              >
                <Image
                  alt={course?.title ?? ""}
                  src={course?.image ?? UOWCAI_COURSE_DEFAULT_IMAGE}
                  style={{
                    opacity: 1,
                    transition: "opacity 0.3s ease-in-out",
                    backgroundColor: "lightgray",
                    height: "auto",
                    width: "100%",
                    minWidth: "360px",
                    maxHeight: "240px",
                    border: "1px solid #f0f0f0",
                    objectFit: "cover",
                  }}
                  preview={false}
                  fallback={UOWCAI_COURSE_DEFAULT_IMAGE}
                  onLoad={() => {
                    dispatch(setCourseDetailImageLoaded(true));
                  }}
                />
              </Flex>
              <Flex
                wrap="wrap"
                gap="middle"
                flex="1 0 60%"
                style={{ margin: "10px" }}
              >
                <div style={{ flex: "1 0 60%", margin: "10px" }}>
                  <Title level={2}>{course?.title}</Title>
                  <Paragraph>
                    <Text>{course?.description}</Text>
                  </Paragraph>
                  <Button
                    type="primary"
                    onClick={() => navigate(course?.learning || "")}
                  >
                    Start
                  </Button>
                </div>
              </Flex>
              <Flex
                wrap="wrap"
                gap="middle"
                flex="1 0 100%"
                style={{ margin: "10px" }}
              >
                <div style={{ flex: "1 0 100%", margin: "10px" }}>
                  <Divider />
                  <Title level={3}>What you will learn</Title>
                </div>
              </Flex>
              <Flex
                wrap="wrap"
                gap="middle"
                flex="1 0 100%"
                style={{ margin: "10px" }}
              >
                {course?.outcomes.map((outcome, index) => (
                  <Flex wrap="wrap" gap="middle" flex="1 0 30%">
                    <Card hoverable type="inner">
                      <Paragraph style={{ minWidth: "100px" }}>
                        <Text italic strong>
                          {outcome}
                        </Text>
                      </Paragraph>
                    </Card>
                  </Flex>
                ))}
              </Flex>
            </Flex>
          </Row>
        </Col>
        <Col span={2} />
      </Row>
    </Layout.Content>
  );
};

export { CourseDetail };
