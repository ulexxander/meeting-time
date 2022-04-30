import { NextPage } from "next";
import Head from "next/head";
import { client } from "../graphql/client";
import {
  TeamByIdDocument,
  TeamByIdQuery,
  TeamByIdQueryVariables,
} from "../graphql/generated";

type HomeProps = { team: TeamByIdQuery["teamByID"] };

export async function getServerSideProps() {
  const { data } = await client.query<TeamByIdQuery, TeamByIdQueryVariables>({
    query: TeamByIdDocument,
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

const Home: NextPage<HomeProps> = ({ team }) => {
  return (
    <div>
      <Head>
        <title>Home page</title>
      </Head>

      <div style={{ textAlign: "center" }}>
        <h1>Hello Next!</h1>

        {team ? (
          <p>
            Team: {team.name} #{team.id}
          </p>
        ) : (
          <p>Team does not exist...</p>
        )}
      </div>
    </div>
  );
};

export default Home;
