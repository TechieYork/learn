import { Layout, Col, Row, List, Image, Typography } from "antd";
import { UOWCAI_COURSES, UOWCAI_COURSE_DEFAULT_IMAGE } from "../../constants/Course";
import { useSelector, useDispatch } from "react-redux";
import { setCoursesImageLoaded } from "./coursesSlice";
import { useNavigate } from "react-router-dom";
const { Title, Link} = Typography;

function Courses() {
  let coursesImageLoaded = useSelector((state: any) => state.courses.coursesImageLoaded);
  let dispatch = useDispatch();
  let navigate = useNavigate();

  return (
    <Layout.Content>
      <Row justify="center">
        <Col span={2} />
        <Col span={20}>
          <Title level={2}>Latest Courses</Title>
          <List
            itemLayout="horizontal"
            dataSource={UOWCAI_COURSES}
            renderItem={(item, index) => (
              <List.Item onClick={() => navigate(item.id)}>
                <List.Item.Meta
                  avatar={
                    <Image
                      alt={item.title}
                      src={item.image}
                      style={{
                        opacity: coursesImageLoaded ? 1 : 0,
                        transition: "opacity 0.3s ease-in-out",
                        backgroundColor: "lightgray",
                        height: "150px",
                        width: "300px",
                        border: "1px solid #f0f0f0",
                        objectFit: "cover",
                      }}
                      preview={false}
                      fallback={UOWCAI_COURSE_DEFAULT_IMAGE}
                      onLoad={() => {
                        dispatch(setCoursesImageLoaded(true));
                        console.log("Image loaded");
                      }}
                    />
                  }
                  title={
                    <Link href={item.url}>
                      <Title level={4}>{item.title}</Title>
                    </Link>
                  }
                  description={item.description}
                />
              </List.Item>
            )}
          />
        </Col>
        <Col span={2} />
      </Row>
    </Layout.Content>
  );
}

export { Courses };
