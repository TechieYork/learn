import React from "react";
import { Flex, Layout, Menu, Image, Space, Avatar, Button, Col } from "antd";
import { UserOutlined } from "@ant-design/icons";
import { useLocation, useNavigate } from "react-router-dom";
import { useSelector } from "react-redux";
import { useDispatch } from "react-redux";
import { setMenu } from "../components/headerSlice";
import { UOW_LOGO } from "../constants/Course";

function Header() {
  let navigate = useNavigate();
  let header = useSelector((state: any) => state.header);
  let dispatch = useDispatch();
  let location = useLocation();

  const defaultMenu =
    location.pathname === "/" ? "home" : location.pathname.substring(1);

  return (
    <Layout.Header
      style={{ display: "flex", boxShadow: "2px 2px 3px #888888", marginBottom: 0 }}
    >
      <Col flex={16} style={{ display: "flex", justifyContent: "flex-start" }}>
        <Space direction="horizontal" size="small">
          <Image
            src={UOW_LOGO}
            alt="UOWCAI"
            preview={false}
            style={{ height: "48px", width: "150px" }}
          ></Image>
          <Menu
            theme="light"
            mode="horizontal"
            style={{ minWidth: "400px" }}
            selectedKeys={[header.selected || defaultMenu]}
            onSelect={(item) => dispatch(setMenu(item.key))}
            onClick={(item) => navigate("/" + item.key)}
          >
            <Menu.Item key="home">Home</Menu.Item>
            <Menu.Item key="courses">Courses</Menu.Item>
            <Menu.Item key="certification">Certification</Menu.Item>
            <Menu.Item key="about" onClick={() => navigate("/about")}>
              About
            </Menu.Item>
          </Menu>
        </Space>
      </Col>
      <Col flex={8} style={{ display: "flex", justifyContent: "flex-end" }}>
        <Space direction="horizontal" size="small">
          <Button type="text">Login</Button>
          <Button type="text">Register</Button>
        </Space>
      </Col>
    </Layout.Header>
  );
}

export { Header };
