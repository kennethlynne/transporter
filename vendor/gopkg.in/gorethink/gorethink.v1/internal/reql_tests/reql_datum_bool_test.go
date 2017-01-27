// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/gorethink/gorethink.v2"
	"gopkg.in/gorethink/gorethink.v2/internal/compare"
)

// Tests of conversion to and from the RQL bool type
func TestDatumBoolSuite(t *testing.T) {
	suite.Run(t, new(DatumBoolSuite))
}

type DatumBoolSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *DatumBoolSuite) SetupTest() {
	suite.T().Log("Setting up DatumBoolSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("test").Exec(suite.session)
	err = r.DBCreate("test").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *DatumBoolSuite) TearDownSuite() {
	suite.T().Log("Tearing down DatumBoolSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *DatumBoolSuite) TestCases() {
	suite.T().Log("Running DatumBoolSuite: Tests of conversion to and from the RQL bool type")

	{
		// datum/bool.yaml line #3
		/* true */
		var expected_ bool = true
		/* r.expr(True) */

		suite.T().Log("About to run line #3: r.Expr(true)")

		runAndAssert(suite.Suite, expected_, r.Expr(true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #3")
	}

	{
		// datum/bool.yaml line #10
		/* false */
		var expected_ bool = false
		/* r.expr(False) */

		suite.T().Log("About to run line #10: r.Expr(false)")

		runAndAssert(suite.Suite, expected_, r.Expr(false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #10")
	}

	{
		// datum/bool.yaml line #17
		/* 'BOOL' */
		var expected_ string = "BOOL"
		/* r.expr(False).type_of() */

		suite.T().Log("About to run line #17: r.Expr(false).TypeOf()")

		runAndAssert(suite.Suite, expected_, r.Expr(false).TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #17")
	}

	{
		// datum/bool.yaml line #21
		/* 'true' */
		var expected_ string = "true"
		/* r.expr(True).coerce_to('string') */

		suite.T().Log("About to run line #21: r.Expr(true).CoerceTo('string')")

		runAndAssert(suite.Suite, expected_, r.Expr(true).CoerceTo("string"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #21")
	}

	{
		// datum/bool.yaml line #24
		/* True */
		var expected_ bool = true
		/* r.expr(True).coerce_to('bool') */

		suite.T().Log("About to run line #24: r.Expr(true).CoerceTo('bool')")

		runAndAssert(suite.Suite, expected_, r.Expr(true).CoerceTo("bool"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #24")
	}

	{
		// datum/bool.yaml line #27
		/* False */
		var expected_ bool = false
		/* r.expr(False).coerce_to('bool') */

		suite.T().Log("About to run line #27: r.Expr(false).CoerceTo('bool')")

		runAndAssert(suite.Suite, expected_, r.Expr(false).CoerceTo("bool"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #27")
	}

	{
		// datum/bool.yaml line #30
		/* False */
		var expected_ bool = false
		/* r.expr(null).coerce_to('bool') */

		suite.T().Log("About to run line #30: r.Expr(nil).CoerceTo('bool')")

		runAndAssert(suite.Suite, expected_, r.Expr(nil).CoerceTo("bool"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #30")
	}

	{
		// datum/bool.yaml line #33
		/* True */
		var expected_ bool = true
		/* r.expr(0).coerce_to('bool') */

		suite.T().Log("About to run line #33: r.Expr(0).CoerceTo('bool')")

		runAndAssert(suite.Suite, expected_, r.Expr(0).CoerceTo("bool"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #33")
	}

	{
		// datum/bool.yaml line #36
		/* True */
		var expected_ bool = true
		/* r.expr('false').coerce_to('bool') */

		suite.T().Log("About to run line #36: r.Expr('false').CoerceTo('bool')")

		runAndAssert(suite.Suite, expected_, r.Expr("false").CoerceTo("bool"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #36")
	}

	{
		// datum/bool.yaml line #39
		/* True */
		var expected_ bool = true
		/* r.expr('foo').coerce_to('bool') */

		suite.T().Log("About to run line #39: r.Expr('foo').CoerceTo('bool')")

		runAndAssert(suite.Suite, expected_, r.Expr("foo").CoerceTo("bool"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #39")
	}

	{
		// datum/bool.yaml line #42
		/* True */
		var expected_ bool = true
		/* r.expr([]).coerce_to('bool') */

		suite.T().Log("About to run line #42: r.Expr([]interface{}{}).CoerceTo('bool')")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{}).CoerceTo("bool"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #42")
	}

	{
		// datum/bool.yaml line #45
		/* True */
		var expected_ bool = true
		/* r.expr({}).coerce_to('bool') */

		suite.T().Log("About to run line #45: r.Expr(map[interface{}]interface{}{}).CoerceTo('bool')")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{}).CoerceTo("bool"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #45")
	}
}
