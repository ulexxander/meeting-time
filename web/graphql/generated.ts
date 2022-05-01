import { gql } from "@apollo/client";
import * as Apollo from "@apollo/client";
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>;
};
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>;
};
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: number;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: any;
  TimeOfDay: string;
};

export type Meeting = {
  __typename?: "Meeting";
  createdAt: Scalars["Time"];
  endedAt: Scalars["Time"];
  id: Scalars["ID"];
  scheduleId: Scalars["Int"];
  startedAt: Scalars["Time"];
  updatedAt?: Maybe<Scalars["Time"]>;
};

export type MeetingCreate = {
  endedAt: Scalars["Time"];
  scheduleId: Scalars["ID"];
  startedAt: Scalars["Time"];
};

export type Mutation = {
  __typename?: "Mutation";
  meetingCreate: Scalars["ID"];
  scheduleCreate: Scalars["ID"];
  teamCreate: Scalars["ID"];
};

export type MutationMeetingCreateArgs = {
  input: MeetingCreate;
};

export type MutationScheduleCreateArgs = {
  input: ScheduleCreate;
};

export type MutationTeamCreateArgs = {
  input: TeamCreate;
};

export type Query = {
  __typename?: "Query";
  meetingByID?: Maybe<Meeting>;
  scheduleByID?: Maybe<Schedule>;
  teamByID?: Maybe<Team>;
};

export type QueryMeetingByIdArgs = {
  id: Scalars["ID"];
};

export type QueryScheduleByIdArgs = {
  id: Scalars["ID"];
};

export type QueryTeamByIdArgs = {
  id: Scalars["ID"];
};

export type Schedule = {
  __typename?: "Schedule";
  createdAt: Scalars["Time"];
  endsAt: Scalars["TimeOfDay"];
  id: Scalars["ID"];
  meetings: Array<Meeting>;
  name: Scalars["String"];
  startsAt: Scalars["TimeOfDay"];
  teamId: Scalars["ID"];
  updatedAt?: Maybe<Scalars["Time"]>;
};

export type ScheduleCreate = {
  endsAt: Scalars["Time"];
  name: Scalars["String"];
  startsAt: Scalars["Time"];
  teamId: Scalars["ID"];
};

export type Team = {
  __typename?: "Team";
  createdAt: Scalars["Time"];
  id: Scalars["ID"];
  name: Scalars["String"];
  schedules: Array<Schedule>;
  updatedAt?: Maybe<Scalars["Time"]>;
};

export type TeamCreate = {
  name: Scalars["String"];
};

export type TeamByIdQueryVariables = Exact<{
  id: Scalars["ID"];
}>;

export type TeamByIdQuery = {
  __typename?: "Query";
  teamByID?: { __typename?: "Team"; id: number; name: string } | null;
};

export type TeamCreateMutationVariables = Exact<{
  input: TeamCreate;
}>;

export type TeamCreateMutation = {
  __typename?: "Mutation";
  teamCreate: number;
};

export const TeamByIdDocument = gql`
  query TeamByID($id: ID!) {
    teamByID(id: $id) {
      id
      name
    }
  }
`;

/**
 * __useTeamByIdQuery__
 *
 * To run a query within a React component, call `useTeamByIdQuery` and pass it any options that fit your needs.
 * When your component renders, `useTeamByIdQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useTeamByIdQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useTeamByIdQuery(
  baseOptions: Apollo.QueryHookOptions<TeamByIdQuery, TeamByIdQueryVariables>
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useQuery<TeamByIdQuery, TeamByIdQueryVariables>(
    TeamByIdDocument,
    options
  );
}
export function useTeamByIdLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    TeamByIdQuery,
    TeamByIdQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useLazyQuery<TeamByIdQuery, TeamByIdQueryVariables>(
    TeamByIdDocument,
    options
  );
}
export type TeamByIdQueryHookResult = ReturnType<typeof useTeamByIdQuery>;
export type TeamByIdLazyQueryHookResult = ReturnType<
  typeof useTeamByIdLazyQuery
>;
export type TeamByIdQueryResult = Apollo.QueryResult<
  TeamByIdQuery,
  TeamByIdQueryVariables
>;
export const TeamCreateDocument = gql`
  mutation TeamCreate($input: TeamCreate!) {
    teamCreate(input: $input)
  }
`;
export type TeamCreateMutationFn = Apollo.MutationFunction<
  TeamCreateMutation,
  TeamCreateMutationVariables
>;

/**
 * __useTeamCreateMutation__
 *
 * To run a mutation, you first call `useTeamCreateMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useTeamCreateMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [teamCreateMutation, { data, loading, error }] = useTeamCreateMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useTeamCreateMutation(
  baseOptions?: Apollo.MutationHookOptions<
    TeamCreateMutation,
    TeamCreateMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useMutation<TeamCreateMutation, TeamCreateMutationVariables>(
    TeamCreateDocument,
    options
  );
}
export type TeamCreateMutationHookResult = ReturnType<
  typeof useTeamCreateMutation
>;
export type TeamCreateMutationResult =
  Apollo.MutationResult<TeamCreateMutation>;
export type TeamCreateMutationOptions = Apollo.BaseMutationOptions<
  TeamCreateMutation,
  TeamCreateMutationVariables
>;
