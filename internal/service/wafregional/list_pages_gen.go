// Code generated by "internal/generate/listpages/main.go -AWSSDKVersion=2 -ListOps=ListActivatedRulesInRuleGroup,ListByteMatchSets,ListGeoMatchSets,ListIPSets,ListRateBasedRules,ListRegexMatchSets,ListRegexPatternSets,ListRules,ListRuleGroups,ListSizeConstraintSets,ListSqlInjectionMatchSets,ListSubscribedRuleGroups,ListWebACLs,ListXssMatchSets -Paginator=NextMarker"; DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func listActivatedRulesInRuleGroupPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListActivatedRulesInRuleGroupInput, fn func(*wafregional.ListActivatedRulesInRuleGroupOutput, bool) bool) error {
	for {
		output, err := conn.ListActivatedRulesInRuleGroup(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listByteMatchSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListByteMatchSetsInput, fn func(*wafregional.ListByteMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListByteMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listGeoMatchSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListGeoMatchSetsInput, fn func(*wafregional.ListGeoMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListGeoMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listIPSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListIPSetsInput, fn func(*wafregional.ListIPSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListIPSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRateBasedRulesPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRateBasedRulesInput, fn func(*wafregional.ListRateBasedRulesOutput, bool) bool) error {
	for {
		output, err := conn.ListRateBasedRules(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRegexMatchSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRegexMatchSetsInput, fn func(*wafregional.ListRegexMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListRegexMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRegexPatternSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRegexPatternSetsInput, fn func(*wafregional.ListRegexPatternSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListRegexPatternSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRuleGroupsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRuleGroupsInput, fn func(*wafregional.ListRuleGroupsOutput, bool) bool) error {
	for {
		output, err := conn.ListRuleGroups(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRulesPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRulesInput, fn func(*wafregional.ListRulesOutput, bool) bool) error {
	for {
		output, err := conn.ListRules(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listSizeConstraintSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListSizeConstraintSetsInput, fn func(*wafregional.ListSizeConstraintSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListSizeConstraintSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listSQLInjectionMatchSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListSqlInjectionMatchSetsInput, fn func(*wafregional.ListSqlInjectionMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListSqlInjectionMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listSubscribedRuleGroupsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListSubscribedRuleGroupsInput, fn func(*wafregional.ListSubscribedRuleGroupsOutput, bool) bool) error {
	for {
		output, err := conn.ListSubscribedRuleGroups(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listWebACLsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListWebACLsInput, fn func(*wafregional.ListWebACLsOutput, bool) bool) error {
	for {
		output, err := conn.ListWebACLs(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listXSSMatchSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListXssMatchSetsInput, fn func(*wafregional.ListXssMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListXssMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
