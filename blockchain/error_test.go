
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:48</date>
//</624461711574831104>

//版权所有（c）2014-2017 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package blockchain

import (
	"testing"
)

//TesterRorCodeStringer测试错误代码类型的字符串化输出。
func TestErrorCodeStringer(t *testing.T) {
	tests := []struct {
		in   ErrorCode
		want string
	}{
		{ErrDuplicateBlock, "ErrDuplicateBlock"},
		{ErrBlockTooBig, "ErrBlockTooBig"},
		{ErrBlockWeightTooHigh, "ErrBlockWeightTooHigh"},
		{ErrBlockVersionTooOld, "ErrBlockVersionTooOld"},
		{ErrInvalidTime, "ErrInvalidTime"},
		{ErrTimeTooOld, "ErrTimeTooOld"},
		{ErrTimeTooNew, "ErrTimeTooNew"},
		{ErrDifficultyTooLow, "ErrDifficultyTooLow"},
		{ErrUnexpectedDifficulty, "ErrUnexpectedDifficulty"},
		{ErrHighHash, "ErrHighHash"},
		{ErrBadMerkleRoot, "ErrBadMerkleRoot"},
		{ErrBadCheckpoint, "ErrBadCheckpoint"},
		{ErrForkTooOld, "ErrForkTooOld"},
		{ErrCheckpointTimeTooOld, "ErrCheckpointTimeTooOld"},
		{ErrNoTransactions, "ErrNoTransactions"},
		{ErrNoTxInputs, "ErrNoTxInputs"},
		{ErrNoTxOutputs, "ErrNoTxOutputs"},
		{ErrTxTooBig, "ErrTxTooBig"},
		{ErrBadTxOutValue, "ErrBadTxOutValue"},
		{ErrDuplicateTxInputs, "ErrDuplicateTxInputs"},
		{ErrBadTxInput, "ErrBadTxInput"},
		{ErrBadCheckpoint, "ErrBadCheckpoint"},
		{ErrMissingTxOut, "ErrMissingTxOut"},
		{ErrUnfinalizedTx, "ErrUnfinalizedTx"},
		{ErrDuplicateTx, "ErrDuplicateTx"},
		{ErrOverwriteTx, "ErrOverwriteTx"},
		{ErrImmatureSpend, "ErrImmatureSpend"},
		{ErrSpendTooHigh, "ErrSpendTooHigh"},
		{ErrBadFees, "ErrBadFees"},
		{ErrTooManySigOps, "ErrTooManySigOps"},
		{ErrFirstTxNotCoinbase, "ErrFirstTxNotCoinbase"},
		{ErrMultipleCoinbases, "ErrMultipleCoinbases"},
		{ErrBadCoinbaseScriptLen, "ErrBadCoinbaseScriptLen"},
		{ErrBadCoinbaseValue, "ErrBadCoinbaseValue"},
		{ErrMissingCoinbaseHeight, "ErrMissingCoinbaseHeight"},
		{ErrBadCoinbaseHeight, "ErrBadCoinbaseHeight"},
		{ErrScriptMalformed, "ErrScriptMalformed"},
		{ErrScriptValidation, "ErrScriptValidation"},
		{ErrUnexpectedWitness, "ErrUnexpectedWitness"},
		{ErrInvalidWitnessCommitment, "ErrInvalidWitnessCommitment"},
		{ErrWitnessCommitmentMismatch, "ErrWitnessCommitmentMismatch"},
		{ErrPreviousBlockUnknown, "ErrPreviousBlockUnknown"},
		{ErrInvalidAncestorBlock, "ErrInvalidAncestorBlock"},
		{ErrPrevBlockNotBest, "ErrPrevBlockNotBest"},
		{0xffff, "Unknown ErrorCode (65535)"},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

//TestRuleError测试RuleError类型的错误输出。
func TestRuleError(t *testing.T) {
	tests := []struct {
		in   RuleError
		want string
	}{
		{
			RuleError{Description: "duplicate block"},
			"duplicate block",
		},
		{
			RuleError{Description: "human-readable error"},
			"human-readable error",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

//TestDeploymentTerror测试DeploymentTerror类型的字符串化输出。
func TestDeploymentError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   DeploymentError
		want string
	}{
		{
			DeploymentError(0),
			"deployment ID 0 does not exist",
		},
		{
			DeploymentError(10),
			"deployment ID 10 does not exist",
		},
		{
			DeploymentError(123),
			"deployment ID 123 does not exist",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}

}

