import React from "react";
import { Row, Col, Layout } from "antd";
import { Home } from "../features/home/Home";
import { Courses } from "../features/courses/Courses";
import { CourseDetail } from "../features/courses/CourseDetail";
import { CourseLearning } from "../features/courses/CourseLearning";
import { Certification } from "../features/certification/Certification";
import { About } from "../features/about/About";
import { NotFound } from "../features/not-found/NotFound";
import { Routes, Route } from "react-router-dom";
import { Header } from "../components/Header";
import { Footer } from "../components/Footer";
import { Space } from "antd";

const Main = () => {
  return (
    <Layout>
      <Header />
      <Row justify="center">
        <Col span={"auto"} />
        <Col flex="1440px">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/home" element={<Home />} />
            <Route path="/courses" element={<Courses />} />
            <Route path="/courses/:courseId" element={<CourseDetail />} />
            <Route path="/learning/:courseId" element={<CourseLearning />} />
            <Route path="/certification" element={<Certification />} />
            <Route path="/about" element={<About />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </Col>
        <Col span={"auto"} />
      </Row>
      <Footer />
    </Layout>
  );
};

export default Main;
