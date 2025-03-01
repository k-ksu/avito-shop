package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/k-ksu/avito-shop/internal/service.TxHistorier -o ./internal/service/mocks/tx_historier.go -n TxHistorier

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/k-ksu/avito-shop/internal/repository/wrapper"
)

// TxHistorier implements service.TxHistorier
type TxHistorier struct {
	t minimock.Tester

	funcAddNew          func(ctx context.Context, tx wrapper.Tx, transaction model.Transaction) (err error)
	inspectFuncAddNew   func(ctx context.Context, tx wrapper.Tx, transaction model.Transaction)
	afterAddNewCounter  uint64
	beforeAddNewCounter uint64
	AddNewMock          mTxHistorierAddNew

	funcGetAllFrom          func(ctx context.Context, tx wrapper.Tx, fromUser int) (sa1 []model.SentCoins, err error)
	inspectFuncGetAllFrom   func(ctx context.Context, tx wrapper.Tx, fromUser int)
	afterGetAllFromCounter  uint64
	beforeGetAllFromCounter uint64
	GetAllFromMock          mTxHistorierGetAllFrom

	funcGetAllTo          func(ctx context.Context, tx wrapper.Tx, toUser int) (ra1 []model.ReceivedCoins, err error)
	inspectFuncGetAllTo   func(ctx context.Context, tx wrapper.Tx, toUser int)
	afterGetAllToCounter  uint64
	beforeGetAllToCounter uint64
	GetAllToMock          mTxHistorierGetAllTo
}

// NewTxHistorier returns a mock for service.TxHistorier
func NewTxHistorier(t minimock.Tester) *TxHistorier {
	m := &TxHistorier{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddNewMock = mTxHistorierAddNew{mock: m}
	m.AddNewMock.callArgs = []*TxHistorierAddNewParams{}

	m.GetAllFromMock = mTxHistorierGetAllFrom{mock: m}
	m.GetAllFromMock.callArgs = []*TxHistorierGetAllFromParams{}

	m.GetAllToMock = mTxHistorierGetAllTo{mock: m}
	m.GetAllToMock.callArgs = []*TxHistorierGetAllToParams{}

	return m
}

type mTxHistorierAddNew struct {
	mock               *TxHistorier
	defaultExpectation *TxHistorierAddNewExpectation
	expectations       []*TxHistorierAddNewExpectation

	callArgs []*TxHistorierAddNewParams
	mutex    sync.RWMutex
}

// TxHistorierAddNewExpectation specifies expectation struct of the TxHistorier.AddNew
type TxHistorierAddNewExpectation struct {
	mock    *TxHistorier
	params  *TxHistorierAddNewParams
	results *TxHistorierAddNewResults
	Counter uint64
}

// TxHistorierAddNewParams contains parameters of the TxHistorier.AddNew
type TxHistorierAddNewParams struct {
	ctx         context.Context
	tx          wrapper.Tx
	transaction model.Transaction
}

// TxHistorierAddNewResults contains results of the TxHistorier.AddNew
type TxHistorierAddNewResults struct {
	err error
}

// Expect sets up expected params for TxHistorier.AddNew
func (mmAddNew *mTxHistorierAddNew) Expect(ctx context.Context, tx wrapper.Tx, transaction model.Transaction) *mTxHistorierAddNew {
	if mmAddNew.mock.funcAddNew != nil {
		mmAddNew.mock.t.Fatalf("TxHistorier.AddNew mock is already set by Set")
	}

	if mmAddNew.defaultExpectation == nil {
		mmAddNew.defaultExpectation = &TxHistorierAddNewExpectation{}
	}

	mmAddNew.defaultExpectation.params = &TxHistorierAddNewParams{ctx, tx, transaction}
	for _, e := range mmAddNew.expectations {
		if minimock.Equal(e.params, mmAddNew.defaultExpectation.params) {
			mmAddNew.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddNew.defaultExpectation.params)
		}
	}

	return mmAddNew
}

// Inspect accepts an inspector function that has same arguments as the TxHistorier.AddNew
func (mmAddNew *mTxHistorierAddNew) Inspect(f func(ctx context.Context, tx wrapper.Tx, transaction model.Transaction)) *mTxHistorierAddNew {
	if mmAddNew.mock.inspectFuncAddNew != nil {
		mmAddNew.mock.t.Fatalf("Inspect function is already set for TxHistorier.AddNew")
	}

	mmAddNew.mock.inspectFuncAddNew = f

	return mmAddNew
}

// Return sets up results that will be returned by TxHistorier.AddNew
func (mmAddNew *mTxHistorierAddNew) Return(err error) *TxHistorier {
	if mmAddNew.mock.funcAddNew != nil {
		mmAddNew.mock.t.Fatalf("TxHistorier.AddNew mock is already set by Set")
	}

	if mmAddNew.defaultExpectation == nil {
		mmAddNew.defaultExpectation = &TxHistorierAddNewExpectation{mock: mmAddNew.mock}
	}
	mmAddNew.defaultExpectation.results = &TxHistorierAddNewResults{err}
	return mmAddNew.mock
}

// Set uses given function f to mock the TxHistorier.AddNew method
func (mmAddNew *mTxHistorierAddNew) Set(f func(ctx context.Context, tx wrapper.Tx, transaction model.Transaction) (err error)) *TxHistorier {
	if mmAddNew.defaultExpectation != nil {
		mmAddNew.mock.t.Fatalf("Default expectation is already set for the TxHistorier.AddNew method")
	}

	if len(mmAddNew.expectations) > 0 {
		mmAddNew.mock.t.Fatalf("Some expectations are already set for the TxHistorier.AddNew method")
	}

	mmAddNew.mock.funcAddNew = f
	return mmAddNew.mock
}

// When sets expectation for the TxHistorier.AddNew which will trigger the result defined by the following
// Then helper
func (mmAddNew *mTxHistorierAddNew) When(ctx context.Context, tx wrapper.Tx, transaction model.Transaction) *TxHistorierAddNewExpectation {
	if mmAddNew.mock.funcAddNew != nil {
		mmAddNew.mock.t.Fatalf("TxHistorier.AddNew mock is already set by Set")
	}

	expectation := &TxHistorierAddNewExpectation{
		mock:   mmAddNew.mock,
		params: &TxHistorierAddNewParams{ctx, tx, transaction},
	}
	mmAddNew.expectations = append(mmAddNew.expectations, expectation)
	return expectation
}

// Then sets up TxHistorier.AddNew return parameters for the expectation previously defined by the When method
func (e *TxHistorierAddNewExpectation) Then(err error) *TxHistorier {
	e.results = &TxHistorierAddNewResults{err}
	return e.mock
}

// AddNew implements service.TxHistorier
func (mmAddNew *TxHistorier) AddNew(ctx context.Context, tx wrapper.Tx, transaction model.Transaction) (err error) {
	mm_atomic.AddUint64(&mmAddNew.beforeAddNewCounter, 1)
	defer mm_atomic.AddUint64(&mmAddNew.afterAddNewCounter, 1)

	if mmAddNew.inspectFuncAddNew != nil {
		mmAddNew.inspectFuncAddNew(ctx, tx, transaction)
	}

	mm_params := &TxHistorierAddNewParams{ctx, tx, transaction}

	// Record call args
	mmAddNew.AddNewMock.mutex.Lock()
	mmAddNew.AddNewMock.callArgs = append(mmAddNew.AddNewMock.callArgs, mm_params)
	mmAddNew.AddNewMock.mutex.Unlock()

	for _, e := range mmAddNew.AddNewMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAddNew.AddNewMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddNew.AddNewMock.defaultExpectation.Counter, 1)
		mm_want := mmAddNew.AddNewMock.defaultExpectation.params
		mm_got := TxHistorierAddNewParams{ctx, tx, transaction}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddNew.t.Errorf("TxHistorier.AddNew got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAddNew.AddNewMock.defaultExpectation.results
		if mm_results == nil {
			mmAddNew.t.Fatal("No results are set for the TxHistorier.AddNew")
		}
		return (*mm_results).err
	}
	if mmAddNew.funcAddNew != nil {
		return mmAddNew.funcAddNew(ctx, tx, transaction)
	}
	mmAddNew.t.Fatalf("Unexpected call to TxHistorier.AddNew. %v %v %v", ctx, tx, transaction)
	return
}

// AddNewAfterCounter returns a count of finished TxHistorier.AddNew invocations
func (mmAddNew *TxHistorier) AddNewAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddNew.afterAddNewCounter)
}

// AddNewBeforeCounter returns a count of TxHistorier.AddNew invocations
func (mmAddNew *TxHistorier) AddNewBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddNew.beforeAddNewCounter)
}

// Calls returns a list of arguments used in each call to TxHistorier.AddNew.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddNew *mTxHistorierAddNew) Calls() []*TxHistorierAddNewParams {
	mmAddNew.mutex.RLock()

	argCopy := make([]*TxHistorierAddNewParams, len(mmAddNew.callArgs))
	copy(argCopy, mmAddNew.callArgs)

	mmAddNew.mutex.RUnlock()

	return argCopy
}

// MinimockAddNewDone returns true if the count of the AddNew invocations corresponds
// the number of defined expectations
func (m *TxHistorier) MinimockAddNewDone() bool {
	for _, e := range m.AddNewMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddNewMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddNewCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddNew != nil && mm_atomic.LoadUint64(&m.afterAddNewCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddNewInspect logs each unmet expectation
func (m *TxHistorier) MinimockAddNewInspect() {
	for _, e := range m.AddNewMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TxHistorier.AddNew with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddNewMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddNewCounter) < 1 {
		if m.AddNewMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TxHistorier.AddNew")
		} else {
			m.t.Errorf("Expected call to TxHistorier.AddNew with params: %#v", *m.AddNewMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddNew != nil && mm_atomic.LoadUint64(&m.afterAddNewCounter) < 1 {
		m.t.Error("Expected call to TxHistorier.AddNew")
	}
}

type mTxHistorierGetAllFrom struct {
	mock               *TxHistorier
	defaultExpectation *TxHistorierGetAllFromExpectation
	expectations       []*TxHistorierGetAllFromExpectation

	callArgs []*TxHistorierGetAllFromParams
	mutex    sync.RWMutex
}

// TxHistorierGetAllFromExpectation specifies expectation struct of the TxHistorier.GetAllFrom
type TxHistorierGetAllFromExpectation struct {
	mock    *TxHistorier
	params  *TxHistorierGetAllFromParams
	results *TxHistorierGetAllFromResults
	Counter uint64
}

// TxHistorierGetAllFromParams contains parameters of the TxHistorier.GetAllFrom
type TxHistorierGetAllFromParams struct {
	ctx      context.Context
	tx       wrapper.Tx
	fromUser int
}

// TxHistorierGetAllFromResults contains results of the TxHistorier.GetAllFrom
type TxHistorierGetAllFromResults struct {
	sa1 []model.SentCoins
	err error
}

// Expect sets up expected params for TxHistorier.GetAllFrom
func (mmGetAllFrom *mTxHistorierGetAllFrom) Expect(ctx context.Context, tx wrapper.Tx, fromUser int) *mTxHistorierGetAllFrom {
	if mmGetAllFrom.mock.funcGetAllFrom != nil {
		mmGetAllFrom.mock.t.Fatalf("TxHistorier.GetAllFrom mock is already set by Set")
	}

	if mmGetAllFrom.defaultExpectation == nil {
		mmGetAllFrom.defaultExpectation = &TxHistorierGetAllFromExpectation{}
	}

	mmGetAllFrom.defaultExpectation.params = &TxHistorierGetAllFromParams{ctx, tx, fromUser}
	for _, e := range mmGetAllFrom.expectations {
		if minimock.Equal(e.params, mmGetAllFrom.defaultExpectation.params) {
			mmGetAllFrom.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetAllFrom.defaultExpectation.params)
		}
	}

	return mmGetAllFrom
}

// Inspect accepts an inspector function that has same arguments as the TxHistorier.GetAllFrom
func (mmGetAllFrom *mTxHistorierGetAllFrom) Inspect(f func(ctx context.Context, tx wrapper.Tx, fromUser int)) *mTxHistorierGetAllFrom {
	if mmGetAllFrom.mock.inspectFuncGetAllFrom != nil {
		mmGetAllFrom.mock.t.Fatalf("Inspect function is already set for TxHistorier.GetAllFrom")
	}

	mmGetAllFrom.mock.inspectFuncGetAllFrom = f

	return mmGetAllFrom
}

// Return sets up results that will be returned by TxHistorier.GetAllFrom
func (mmGetAllFrom *mTxHistorierGetAllFrom) Return(sa1 []model.SentCoins, err error) *TxHistorier {
	if mmGetAllFrom.mock.funcGetAllFrom != nil {
		mmGetAllFrom.mock.t.Fatalf("TxHistorier.GetAllFrom mock is already set by Set")
	}

	if mmGetAllFrom.defaultExpectation == nil {
		mmGetAllFrom.defaultExpectation = &TxHistorierGetAllFromExpectation{mock: mmGetAllFrom.mock}
	}
	mmGetAllFrom.defaultExpectation.results = &TxHistorierGetAllFromResults{sa1, err}
	return mmGetAllFrom.mock
}

// Set uses given function f to mock the TxHistorier.GetAllFrom method
func (mmGetAllFrom *mTxHistorierGetAllFrom) Set(f func(ctx context.Context, tx wrapper.Tx, fromUser int) (sa1 []model.SentCoins, err error)) *TxHistorier {
	if mmGetAllFrom.defaultExpectation != nil {
		mmGetAllFrom.mock.t.Fatalf("Default expectation is already set for the TxHistorier.GetAllFrom method")
	}

	if len(mmGetAllFrom.expectations) > 0 {
		mmGetAllFrom.mock.t.Fatalf("Some expectations are already set for the TxHistorier.GetAllFrom method")
	}

	mmGetAllFrom.mock.funcGetAllFrom = f
	return mmGetAllFrom.mock
}

// When sets expectation for the TxHistorier.GetAllFrom which will trigger the result defined by the following
// Then helper
func (mmGetAllFrom *mTxHistorierGetAllFrom) When(ctx context.Context, tx wrapper.Tx, fromUser int) *TxHistorierGetAllFromExpectation {
	if mmGetAllFrom.mock.funcGetAllFrom != nil {
		mmGetAllFrom.mock.t.Fatalf("TxHistorier.GetAllFrom mock is already set by Set")
	}

	expectation := &TxHistorierGetAllFromExpectation{
		mock:   mmGetAllFrom.mock,
		params: &TxHistorierGetAllFromParams{ctx, tx, fromUser},
	}
	mmGetAllFrom.expectations = append(mmGetAllFrom.expectations, expectation)
	return expectation
}

// Then sets up TxHistorier.GetAllFrom return parameters for the expectation previously defined by the When method
func (e *TxHistorierGetAllFromExpectation) Then(sa1 []model.SentCoins, err error) *TxHistorier {
	e.results = &TxHistorierGetAllFromResults{sa1, err}
	return e.mock
}

// GetAllFrom implements service.TxHistorier
func (mmGetAllFrom *TxHistorier) GetAllFrom(ctx context.Context, tx wrapper.Tx, fromUser int) (sa1 []model.SentCoins, err error) {
	mm_atomic.AddUint64(&mmGetAllFrom.beforeGetAllFromCounter, 1)
	defer mm_atomic.AddUint64(&mmGetAllFrom.afterGetAllFromCounter, 1)

	if mmGetAllFrom.inspectFuncGetAllFrom != nil {
		mmGetAllFrom.inspectFuncGetAllFrom(ctx, tx, fromUser)
	}

	mm_params := &TxHistorierGetAllFromParams{ctx, tx, fromUser}

	// Record call args
	mmGetAllFrom.GetAllFromMock.mutex.Lock()
	mmGetAllFrom.GetAllFromMock.callArgs = append(mmGetAllFrom.GetAllFromMock.callArgs, mm_params)
	mmGetAllFrom.GetAllFromMock.mutex.Unlock()

	for _, e := range mmGetAllFrom.GetAllFromMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.sa1, e.results.err
		}
	}

	if mmGetAllFrom.GetAllFromMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetAllFrom.GetAllFromMock.defaultExpectation.Counter, 1)
		mm_want := mmGetAllFrom.GetAllFromMock.defaultExpectation.params
		mm_got := TxHistorierGetAllFromParams{ctx, tx, fromUser}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetAllFrom.t.Errorf("TxHistorier.GetAllFrom got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetAllFrom.GetAllFromMock.defaultExpectation.results
		if mm_results == nil {
			mmGetAllFrom.t.Fatal("No results are set for the TxHistorier.GetAllFrom")
		}
		return (*mm_results).sa1, (*mm_results).err
	}
	if mmGetAllFrom.funcGetAllFrom != nil {
		return mmGetAllFrom.funcGetAllFrom(ctx, tx, fromUser)
	}
	mmGetAllFrom.t.Fatalf("Unexpected call to TxHistorier.GetAllFrom. %v %v %v", ctx, tx, fromUser)
	return
}

// GetAllFromAfterCounter returns a count of finished TxHistorier.GetAllFrom invocations
func (mmGetAllFrom *TxHistorier) GetAllFromAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetAllFrom.afterGetAllFromCounter)
}

// GetAllFromBeforeCounter returns a count of TxHistorier.GetAllFrom invocations
func (mmGetAllFrom *TxHistorier) GetAllFromBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetAllFrom.beforeGetAllFromCounter)
}

// Calls returns a list of arguments used in each call to TxHistorier.GetAllFrom.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetAllFrom *mTxHistorierGetAllFrom) Calls() []*TxHistorierGetAllFromParams {
	mmGetAllFrom.mutex.RLock()

	argCopy := make([]*TxHistorierGetAllFromParams, len(mmGetAllFrom.callArgs))
	copy(argCopy, mmGetAllFrom.callArgs)

	mmGetAllFrom.mutex.RUnlock()

	return argCopy
}

// MinimockGetAllFromDone returns true if the count of the GetAllFrom invocations corresponds
// the number of defined expectations
func (m *TxHistorier) MinimockGetAllFromDone() bool {
	for _, e := range m.GetAllFromMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetAllFromMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetAllFromCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetAllFrom != nil && mm_atomic.LoadUint64(&m.afterGetAllFromCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetAllFromInspect logs each unmet expectation
func (m *TxHistorier) MinimockGetAllFromInspect() {
	for _, e := range m.GetAllFromMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TxHistorier.GetAllFrom with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetAllFromMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetAllFromCounter) < 1 {
		if m.GetAllFromMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TxHistorier.GetAllFrom")
		} else {
			m.t.Errorf("Expected call to TxHistorier.GetAllFrom with params: %#v", *m.GetAllFromMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetAllFrom != nil && mm_atomic.LoadUint64(&m.afterGetAllFromCounter) < 1 {
		m.t.Error("Expected call to TxHistorier.GetAllFrom")
	}
}

type mTxHistorierGetAllTo struct {
	mock               *TxHistorier
	defaultExpectation *TxHistorierGetAllToExpectation
	expectations       []*TxHistorierGetAllToExpectation

	callArgs []*TxHistorierGetAllToParams
	mutex    sync.RWMutex
}

// TxHistorierGetAllToExpectation specifies expectation struct of the TxHistorier.GetAllTo
type TxHistorierGetAllToExpectation struct {
	mock    *TxHistorier
	params  *TxHistorierGetAllToParams
	results *TxHistorierGetAllToResults
	Counter uint64
}

// TxHistorierGetAllToParams contains parameters of the TxHistorier.GetAllTo
type TxHistorierGetAllToParams struct {
	ctx    context.Context
	tx     wrapper.Tx
	toUser int
}

// TxHistorierGetAllToResults contains results of the TxHistorier.GetAllTo
type TxHistorierGetAllToResults struct {
	ra1 []model.ReceivedCoins
	err error
}

// Expect sets up expected params for TxHistorier.GetAllTo
func (mmGetAllTo *mTxHistorierGetAllTo) Expect(ctx context.Context, tx wrapper.Tx, toUser int) *mTxHistorierGetAllTo {
	if mmGetAllTo.mock.funcGetAllTo != nil {
		mmGetAllTo.mock.t.Fatalf("TxHistorier.GetAllTo mock is already set by Set")
	}

	if mmGetAllTo.defaultExpectation == nil {
		mmGetAllTo.defaultExpectation = &TxHistorierGetAllToExpectation{}
	}

	mmGetAllTo.defaultExpectation.params = &TxHistorierGetAllToParams{ctx, tx, toUser}
	for _, e := range mmGetAllTo.expectations {
		if minimock.Equal(e.params, mmGetAllTo.defaultExpectation.params) {
			mmGetAllTo.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetAllTo.defaultExpectation.params)
		}
	}

	return mmGetAllTo
}

// Inspect accepts an inspector function that has same arguments as the TxHistorier.GetAllTo
func (mmGetAllTo *mTxHistorierGetAllTo) Inspect(f func(ctx context.Context, tx wrapper.Tx, toUser int)) *mTxHistorierGetAllTo {
	if mmGetAllTo.mock.inspectFuncGetAllTo != nil {
		mmGetAllTo.mock.t.Fatalf("Inspect function is already set for TxHistorier.GetAllTo")
	}

	mmGetAllTo.mock.inspectFuncGetAllTo = f

	return mmGetAllTo
}

// Return sets up results that will be returned by TxHistorier.GetAllTo
func (mmGetAllTo *mTxHistorierGetAllTo) Return(ra1 []model.ReceivedCoins, err error) *TxHistorier {
	if mmGetAllTo.mock.funcGetAllTo != nil {
		mmGetAllTo.mock.t.Fatalf("TxHistorier.GetAllTo mock is already set by Set")
	}

	if mmGetAllTo.defaultExpectation == nil {
		mmGetAllTo.defaultExpectation = &TxHistorierGetAllToExpectation{mock: mmGetAllTo.mock}
	}
	mmGetAllTo.defaultExpectation.results = &TxHistorierGetAllToResults{ra1, err}
	return mmGetAllTo.mock
}

// Set uses given function f to mock the TxHistorier.GetAllTo method
func (mmGetAllTo *mTxHistorierGetAllTo) Set(f func(ctx context.Context, tx wrapper.Tx, toUser int) (ra1 []model.ReceivedCoins, err error)) *TxHistorier {
	if mmGetAllTo.defaultExpectation != nil {
		mmGetAllTo.mock.t.Fatalf("Default expectation is already set for the TxHistorier.GetAllTo method")
	}

	if len(mmGetAllTo.expectations) > 0 {
		mmGetAllTo.mock.t.Fatalf("Some expectations are already set for the TxHistorier.GetAllTo method")
	}

	mmGetAllTo.mock.funcGetAllTo = f
	return mmGetAllTo.mock
}

// When sets expectation for the TxHistorier.GetAllTo which will trigger the result defined by the following
// Then helper
func (mmGetAllTo *mTxHistorierGetAllTo) When(ctx context.Context, tx wrapper.Tx, toUser int) *TxHistorierGetAllToExpectation {
	if mmGetAllTo.mock.funcGetAllTo != nil {
		mmGetAllTo.mock.t.Fatalf("TxHistorier.GetAllTo mock is already set by Set")
	}

	expectation := &TxHistorierGetAllToExpectation{
		mock:   mmGetAllTo.mock,
		params: &TxHistorierGetAllToParams{ctx, tx, toUser},
	}
	mmGetAllTo.expectations = append(mmGetAllTo.expectations, expectation)
	return expectation
}

// Then sets up TxHistorier.GetAllTo return parameters for the expectation previously defined by the When method
func (e *TxHistorierGetAllToExpectation) Then(ra1 []model.ReceivedCoins, err error) *TxHistorier {
	e.results = &TxHistorierGetAllToResults{ra1, err}
	return e.mock
}

// GetAllTo implements service.TxHistorier
func (mmGetAllTo *TxHistorier) GetAllTo(ctx context.Context, tx wrapper.Tx, toUser int) (ra1 []model.ReceivedCoins, err error) {
	mm_atomic.AddUint64(&mmGetAllTo.beforeGetAllToCounter, 1)
	defer mm_atomic.AddUint64(&mmGetAllTo.afterGetAllToCounter, 1)

	if mmGetAllTo.inspectFuncGetAllTo != nil {
		mmGetAllTo.inspectFuncGetAllTo(ctx, tx, toUser)
	}

	mm_params := &TxHistorierGetAllToParams{ctx, tx, toUser}

	// Record call args
	mmGetAllTo.GetAllToMock.mutex.Lock()
	mmGetAllTo.GetAllToMock.callArgs = append(mmGetAllTo.GetAllToMock.callArgs, mm_params)
	mmGetAllTo.GetAllToMock.mutex.Unlock()

	for _, e := range mmGetAllTo.GetAllToMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ra1, e.results.err
		}
	}

	if mmGetAllTo.GetAllToMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetAllTo.GetAllToMock.defaultExpectation.Counter, 1)
		mm_want := mmGetAllTo.GetAllToMock.defaultExpectation.params
		mm_got := TxHistorierGetAllToParams{ctx, tx, toUser}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetAllTo.t.Errorf("TxHistorier.GetAllTo got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetAllTo.GetAllToMock.defaultExpectation.results
		if mm_results == nil {
			mmGetAllTo.t.Fatal("No results are set for the TxHistorier.GetAllTo")
		}
		return (*mm_results).ra1, (*mm_results).err
	}
	if mmGetAllTo.funcGetAllTo != nil {
		return mmGetAllTo.funcGetAllTo(ctx, tx, toUser)
	}
	mmGetAllTo.t.Fatalf("Unexpected call to TxHistorier.GetAllTo. %v %v %v", ctx, tx, toUser)
	return
}

// GetAllToAfterCounter returns a count of finished TxHistorier.GetAllTo invocations
func (mmGetAllTo *TxHistorier) GetAllToAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetAllTo.afterGetAllToCounter)
}

// GetAllToBeforeCounter returns a count of TxHistorier.GetAllTo invocations
func (mmGetAllTo *TxHistorier) GetAllToBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetAllTo.beforeGetAllToCounter)
}

// Calls returns a list of arguments used in each call to TxHistorier.GetAllTo.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetAllTo *mTxHistorierGetAllTo) Calls() []*TxHistorierGetAllToParams {
	mmGetAllTo.mutex.RLock()

	argCopy := make([]*TxHistorierGetAllToParams, len(mmGetAllTo.callArgs))
	copy(argCopy, mmGetAllTo.callArgs)

	mmGetAllTo.mutex.RUnlock()

	return argCopy
}

// MinimockGetAllToDone returns true if the count of the GetAllTo invocations corresponds
// the number of defined expectations
func (m *TxHistorier) MinimockGetAllToDone() bool {
	for _, e := range m.GetAllToMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetAllToMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetAllToCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetAllTo != nil && mm_atomic.LoadUint64(&m.afterGetAllToCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetAllToInspect logs each unmet expectation
func (m *TxHistorier) MinimockGetAllToInspect() {
	for _, e := range m.GetAllToMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TxHistorier.GetAllTo with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetAllToMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetAllToCounter) < 1 {
		if m.GetAllToMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TxHistorier.GetAllTo")
		} else {
			m.t.Errorf("Expected call to TxHistorier.GetAllTo with params: %#v", *m.GetAllToMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetAllTo != nil && mm_atomic.LoadUint64(&m.afterGetAllToCounter) < 1 {
		m.t.Error("Expected call to TxHistorier.GetAllTo")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TxHistorier) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddNewInspect()

		m.MinimockGetAllFromInspect()

		m.MinimockGetAllToInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TxHistorier) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *TxHistorier) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddNewDone() &&
		m.MinimockGetAllFromDone() &&
		m.MinimockGetAllToDone()
}
