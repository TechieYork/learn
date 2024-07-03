import { Layout, Breadcrumb, Menu, theme } from "antd";
import { useParams } from "react-router-dom";
import { Pdf } from "../../components/Pdf";
import { UOWCAI_COURSES, UOWCAI_COURSES_OUTLINE } from "../../constants/Course";
import { useSelector, useDispatch } from "react-redux";
import { setItemKey, setSection, setTopic } from "./courseLearningSlice";
const { Content, Sider } = Layout;

const CourseLearning = () => {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  // get courseID from URL
  const params = useParams();
  const courseID = params.courseId;

  // get course outline from storage
  const courseOutline =
    UOWCAI_COURSES_OUTLINE[courseID as keyof typeof UOWCAI_COURSES_OUTLINE];

  // get topic state
  const itemKey = useSelector((state: any) => state.courseLearning.itemKey);
  const section = useSelector((state: any) => state.courseLearning.section);
  const topic = useSelector((state: any) => state.courseLearning.topic);

  const dispatch = useDispatch();

  return (
    <Layout.Content style={{ zIndex: 100, margin: "3px" }}>
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
                  const subKey = j + 1;
                  return {
                    key: `${key}-${subKey}-${topic.page}`,
                    label: `${key}.${subKey} ${topic.name}`,
                    // style: { whiteSpace: "normal", height: "auto" },
                  };
                }),
              };
            })}
            onSelect={(item) =>
              dispatch(setItemKey(item.key)) &&
              dispatch(
                setSection(
                  courseOutline.outline[parseInt(item.key.split("-")[0]) - 1]
                    .name
                )
              ) &&
              dispatch(
                setTopic(
                  courseOutline.outline[parseInt(item.key.split("-")[0]) - 1]
                    .topics[parseInt(item.key.split("-")[1]) - 1].name
                )
              )
            }
          />
        </Sider>
        <Content style={{ padding: "24px 96px", minHeight: 280 }}>
          <Breadcrumb style={{ margin: "16px 0" }}>
            <Breadcrumb.Item>
              {UOWCAI_COURSES.find((element) => element.id === courseID)
                ?.title || ""}
            </Breadcrumb.Item>
            <Breadcrumb.Item>{section}</Breadcrumb.Item>
            <Breadcrumb.Item>{topic}</Breadcrumb.Item>
          </Breadcrumb>
          <Pdf
            file={
              UOWCAI_COURSES.find((element) => element.id === courseID)?.pdf ||
              ""
            }
            page={parseInt(itemKey ? itemKey.split("-")[2] : "1-1-1")}
          />
        </Content>
      </Layout>
    </Layout.Content>
  );
};

export { CourseLearning };
