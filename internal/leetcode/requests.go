package leetcode

func BuildLeetcodeRecentStatsQuery(leetcodeId string) (string, map[string]interface{}) {
	query := `
		recentAcSubmissions($username: STring!, $limit: Int!) {
			recentAcSubmissions(username: $username, limit: $limit) {
				id
				title
				titleSlug
				timestamp
			}
		}
	`

	variables := map[string]interface{}{
		"username": leetcodeId,
		"limit":    dataLimit,
	}

	return query, variables
}

func BuildUserRankQuery(leetcodeId string) (string, map[string]string) {
	query := `
		query getUserProfile($username: String!) { 
            allQuestionsCount { 
                difficulty count 
            } matchedUser(username: $username) { 
                 contributions { 
                     points 
                } profile { 
                     reputation ranking 
			    } submissionCalendar submitStats { 
                     acSubmissionNum { 
                        difficulty count submissions 
                    } totalSubmissionNum { 
                         difficulty count submissions 
                       } 
                 } 
            } 
       }
	`

	variables := map[string]string{
		"username": leetcodeId,
	}

	return query, variables
}

func BuildUserSubmissionsQuery(leetcodeId string, offset int, skip int) (string, map[string]interface{}) {
	query := `
		query userSolutionTopics($username: String!, $orderBy: TopicSortingOption, $skip: Int, $first: Int) {
			userSolutionTopics(
				username: $username
				orderBy: $orderBy
				skip: $skip
				first: $first
			) {
				pageInfo {
					hasNextPage
				}
				edges {
					node {
						id
						title
						url
						viewCount
						questionTitle
						post {
							creationDate
							voteCount
						}
					}
				}
			}
		}
	`

	variables := map[string]interface{}{
		"usernmae": leetcodeId,
		"orderBy":  "newest_to_oldest",
		"skip":     skip,
		"first":    offset,
	}

	return query, variables
}

func BuildSolutionByIdRequest(solutionId int) (string, map[string]int) {
	query := `
		query communitySolution($topicId: Int!) {
			isSolutionTopic(id: $topicId)
			topic(id: $topicId) {
				id
				viewCount
				topLevelCommentCount
				favoriteCount
				subscribed
				title
				pinned
				solutionTags {
					name
					slug
				}
				hideFromTrending
				commentCount
				isFavorite
				post {
					id
					voteCount
					voteStatus
					content
					updationDate
					creationDate
					status
					isHidden
				author {
					isDiscussAdmin
					isDiscussStaff
					username
					nameColor
					activeBadge {
						displayName
						icon
					}
					profile {
						userAvatar
						reputation
					}
					isActive
				}
				authorIsModerator
				isOwnPost
			}
		}
	}
	`

	variables := map[string]int{
		"topicId": solutionId,
	}

	return query, variables
}
