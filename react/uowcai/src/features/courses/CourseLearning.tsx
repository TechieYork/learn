import React, { Children } from "react";
import { Flex, Layout, Breadcrumb, Menu, MenuProps, theme } from "antd";
import {
  LaptopOutlined,
  NotificationOutlined,
  UserOutlined,
} from "@ant-design/icons";
import { useParams, useNavigate, useSearchParams } from "react-router-dom";
import { Pdf } from "../../components/Pdf";
import { UOWCAI_COURSES, UOWCAI_COURSES_OUTLINE } from "../../constants/Course";

const { Content, Sider } = Layout;

// render UOWCAI_COURSES_OUTLINE
// const items: MenuProps["items"] = UOWCAI_COURSES_OUTLINE.gai.outline.map(
//   (section, index) => {
//     const key = String(index + 1);

//     return {
//       key: `${key}-${section.start}`,
//       label: `${key}. ${section.name}`,

//       children: section.topics.map((topic, j) => {
//         const subKey = index * 4 + j + 1;
//         return {
//           key: `${key}-${subKey}-${topic.page}`,
//           label: `${key}.${subKey} ${topic.name}`,
//         };
//       }),
//     };
//   }
// );

const CourseLearning = () => {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  const params = useParams();
  const [query] = useSearchParams();
  const courseID = params.courseId;
  const courseOutline =
    UOWCAI_COURSES_OUTLINE[courseID as keyof typeof UOWCAI_COURSES_OUTLINE];
  const navigate = useNavigate();

  return (
    <Layout.Content style={{ zIndex: 100, margin: "3px" }}>
      {/* <Breadcrumb style={{ margin: "16px 0" }}>
          <Breadcrumb.Item>Home</Breadcrumb.Item>
          <Breadcrumb.Item>List</Breadcrumb.Item>
          <Breadcrumb.Item>App</Breadcrumb.Item>
        </Breadcrumb> */}
      <Layout
        style={{
          background: colorBgContainer,
          borderRadius: borderRadiusLG,
          minHeight: "100vh",
        }}
      >
        <Sider style={{ background: colorBgContainer }} width={260}>
          <Menu
            theme="light"
            mode="inline"
            inlineCollapsed={false}
            // defaultSelectedKeys={["1"]}
            defaultOpenKeys={courseOutline.outline.map((section, index) => {
              const key = String(index + 1);
              return `${key}-${section.start}`;
            })}
            style={{ height: "100%", width: "auto" }}
            items={courseOutline.outline.map((section, index) => {
              const key = String(index + 1);

              return {
                key: `${key}-${section.start}`,
                label: `${key}. ${section.name}`,

                children: section.topics.map((topic, j) => {
                  const subKey = index * 4 + j + 1;
                  return {
                    key: `${key}-${subKey}-${topic.page}`,
                    label: `${key}.${subKey} ${topic.name}`,
                  };
                }),
              };
            })}
            onSelect={(item) =>
              navigate(`/learning/${courseID}?page=${item.key.split("-")[2]}`)
            }
          />
        </Sider>
        <Content style={{ padding: "24px 48px", minHeight: 280 }}>
          <Pdf
            file={
              UOWCAI_COURSES.find((element) => element.id === courseID)?.pdf ||
              ""
            }
            page={parseInt(query.get("page") || "1")}
          />
        </Content>
      </Layout>
    </Layout.Content>
  );
};

export { CourseLearning };
