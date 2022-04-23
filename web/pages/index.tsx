import { gql } from "@apollo/client";
import { NextPage } from "next";
import Head from "next/head";
import { client } from "../graphql/client";

const query = gql`
  query TeamByID($id: ID!) {
    teamByID(id: $id) {
      id
      name
    }
  }
`;

export async function getServerSideProps() {
  const { data } = await client.query({
    query,
    variables: {
      id: "1",
    },
  });

  return {
    props: {
      team: data.teamByID,
    },
  };
}

type Team = {
  id: number;
  name: string;
};

const Home: NextPage<{ team: Team }> = ({ team }) => {
  return (
    <div>
      <Head>
        <title>Home page</title>
      </Head>

      <div style={{ textAlign: "center" }}>
        <h1>Hello Next!</h1>

        <p>
          Team: {team.name} #{team.id}
        </p>
      </div>
    </div>
  );
};

export default Home;
