import { NextPage } from "next";
import Head from "next/head";

const Home: NextPage = () => {
  return (
    <div>
      <Head>
        <title>Home page</title>
      </Head>

      <p style={{ textAlign: "center" }}>Hello Next!</p>
    </div>
  );
};

export default Home;
