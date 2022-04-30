import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: any;
};

export type Meeting = {
  __typename?: 'Meeting';
  createdAt: Scalars['Time'];
  endedAt: Scalars['Time'];
  id: Scalars['ID'];
  scheduleId: Scalars['Int'];
  startedAt: Scalars['Time'];
  updatedAt?: Maybe<Scalars['Time']>;
};

export type MeetingCreate = {
  endedAt: Scalars['Time'];
  scheduleId: Scalars['ID'];
  startedAt: Scalars['Time'];
};

export type Mutation = {
  __typename?: 'Mutation';
  meetingCreate: Scalars['ID'];
  scheduleCreate: Scalars['ID'];
  teamCreate: Scalars['ID'];
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
  __typename?: 'Query';
  meetingByID?: Maybe<Meeting>;
  scheduleByID?: Maybe<Schedule>;
  teamByID?: Maybe<Team>;
};


export type QueryMeetingByIdArgs = {
  id: Scalars['ID'];
};


export type QueryScheduleByIdArgs = {
  id: Scalars['ID'];
};


export type QueryTeamByIdArgs = {
  id: Scalars['ID'];
};

export type Schedule = {
  __typename?: 'Schedule';
  createdAt: Scalars['Time'];
  endsAt: Scalars['Time'];
  id: Scalars['ID'];
  meetings: Array<Meeting>;
  name: Scalars['String'];
  startsAt: Scalars['Time'];
  teamId: Scalars['ID'];
  updatedAt?: Maybe<Scalars['Time']>;
};

export type ScheduleCreate = {
  endsAt: Scalars['Time'];
  name: Scalars['String'];
  startsAt: Scalars['Time'];
  teamId: Scalars['ID'];
};

export type Team = {
  __typename?: 'Team';
  createdAt: Scalars['Time'];
  id: Scalars['ID'];
  name: Scalars['String'];
  schedules: Array<Schedule>;
  updatedAt?: Maybe<Scalars['Time']>;
};

export type TeamCreate = {
  name: Scalars['String'];
};

export type TeamByIdQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type TeamByIdQuery = { __typename?: 'Query', teamByID?: { __typename?: 'Team', id: string, name: string } | null };


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
export function useTeamByIdQuery(baseOptions: Apollo.QueryHookOptions<TeamByIdQuery, TeamByIdQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<TeamByIdQuery, TeamByIdQueryVariables>(TeamByIdDocument, options);
      }
export function useTeamByIdLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<TeamByIdQuery, TeamByIdQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<TeamByIdQuery, TeamByIdQueryVariables>(TeamByIdDocument, options);
        }
export type TeamByIdQueryHookResult = ReturnType<typeof useTeamByIdQuery>;
export type TeamByIdLazyQueryHookResult = ReturnType<typeof useTeamByIdLazyQuery>;
export type TeamByIdQueryResult = Apollo.QueryResult<TeamByIdQuery, TeamByIdQueryVariables>;