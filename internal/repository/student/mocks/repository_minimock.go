package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/olezhek28/system-design-party-bot/internal/repository/student.Repository -o ./mocks/repository_minimock.go -n RepositoryMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	mm_student_repository "github.com/olezhek28/system-design-party-bot/internal/repository/student"
)

// RepositoryMock implements student_repository.Repository
type RepositoryMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, student *model.Student) (err error)
	inspectFuncCreate   func(ctx context.Context, student *model.Student)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mRepositoryMockCreate

	funcGetList          func(ctx context.Context, filter *mm_student_repository.Query) (spa1 []*model.Student, err error)
	inspectFuncGetList   func(ctx context.Context, filter *mm_student_repository.Query)
	afterGetListCounter  uint64
	beforeGetListCounter uint64
	GetListMock          mRepositoryMockGetList

	funcIsExist          func(ctx context.Context, telegramChatID int64) (b1 bool, err error)
	inspectFuncIsExist   func(ctx context.Context, telegramChatID int64)
	afterIsExistCounter  uint64
	beforeIsExistCounter uint64
	IsExistMock          mRepositoryMockIsExist

	funcUpdate          func(ctx context.Context, telegramID int64, updateStudent *model.UpdateStudent) (err error)
	inspectFuncUpdate   func(ctx context.Context, telegramID int64, updateStudent *model.UpdateStudent)
	afterUpdateCounter  uint64
	beforeUpdateCounter uint64
	UpdateMock          mRepositoryMockUpdate
}

// NewRepositoryMock returns a mock for student_repository.Repository
func NewRepositoryMock(t minimock.Tester) *RepositoryMock {
	m := &RepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*RepositoryMockCreateParams{}

	m.GetListMock = mRepositoryMockGetList{mock: m}
	m.GetListMock.callArgs = []*RepositoryMockGetListParams{}

	m.IsExistMock = mRepositoryMockIsExist{mock: m}
	m.IsExistMock.callArgs = []*RepositoryMockIsExistParams{}

	m.UpdateMock = mRepositoryMockUpdate{mock: m}
	m.UpdateMock.callArgs = []*RepositoryMockUpdateParams{}

	return m
}

type mRepositoryMockCreate struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockCreateExpectation
	expectations       []*RepositoryMockCreateExpectation

	callArgs []*RepositoryMockCreateParams
	mutex    sync.RWMutex
}

// RepositoryMockCreateExpectation specifies expectation struct of the Repository.Create
type RepositoryMockCreateExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockCreateParams
	results *RepositoryMockCreateResults
	Counter uint64
}

// RepositoryMockCreateParams contains parameters of the Repository.Create
type RepositoryMockCreateParams struct {
	ctx     context.Context
	student *model.Student
}

// RepositoryMockCreateResults contains results of the Repository.Create
type RepositoryMockCreateResults struct {
	err error
}

// Expect sets up expected params for Repository.Create
func (mmCreate *mRepositoryMockCreate) Expect(ctx context.Context, student *model.Student) *mRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("RepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &RepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &RepositoryMockCreateParams{ctx, student}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the Repository.Create
func (mmCreate *mRepositoryMockCreate) Inspect(f func(ctx context.Context, student *model.Student)) *mRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for RepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by Repository.Create
func (mmCreate *mRepositoryMockCreate) Return(err error) *RepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("RepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &RepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &RepositoryMockCreateResults{err}
	return mmCreate.mock
}

//Set uses given function f to mock the Repository.Create method
func (mmCreate *mRepositoryMockCreate) Set(f func(ctx context.Context, student *model.Student) (err error)) *RepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the Repository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the Repository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the Repository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mRepositoryMockCreate) When(ctx context.Context, student *model.Student) *RepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("RepositoryMock.Create mock is already set by Set")
	}

	expectation := &RepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &RepositoryMockCreateParams{ctx, student},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up Repository.Create return parameters for the expectation previously defined by the When method
func (e *RepositoryMockCreateExpectation) Then(err error) *RepositoryMock {
	e.results = &RepositoryMockCreateResults{err}
	return e.mock
}

// Create implements student_repository.Repository
func (mmCreate *RepositoryMock) Create(ctx context.Context, student *model.Student) (err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, student)
	}

	mm_params := &RepositoryMockCreateParams{ctx, student}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := RepositoryMockCreateParams{ctx, student}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("RepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the RepositoryMock.Create")
		}
		return (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, student)
	}
	mmCreate.t.Fatalf("Unexpected call to RepositoryMock.Create. %v %v", ctx, student)
	return
}

// CreateAfterCounter returns a count of finished RepositoryMock.Create invocations
func (mmCreate *RepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of RepositoryMock.Create invocations
func (mmCreate *RepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mRepositoryMockCreate) Calls() []*RepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*RepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *RepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.Create")
	}
}

type mRepositoryMockGetList struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockGetListExpectation
	expectations       []*RepositoryMockGetListExpectation

	callArgs []*RepositoryMockGetListParams
	mutex    sync.RWMutex
}

// RepositoryMockGetListExpectation specifies expectation struct of the Repository.GetList
type RepositoryMockGetListExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockGetListParams
	results *RepositoryMockGetListResults
	Counter uint64
}

// RepositoryMockGetListParams contains parameters of the Repository.GetList
type RepositoryMockGetListParams struct {
	ctx    context.Context
	filter *mm_student_repository.Query
}

// RepositoryMockGetListResults contains results of the Repository.GetList
type RepositoryMockGetListResults struct {
	spa1 []*model.Student
	err  error
}

// Expect sets up expected params for Repository.GetList
func (mmGetList *mRepositoryMockGetList) Expect(ctx context.Context, filter *mm_student_repository.Query) *mRepositoryMockGetList {
	if mmGetList.mock.funcGetList != nil {
		mmGetList.mock.t.Fatalf("RepositoryMock.GetList mock is already set by Set")
	}

	if mmGetList.defaultExpectation == nil {
		mmGetList.defaultExpectation = &RepositoryMockGetListExpectation{}
	}

	mmGetList.defaultExpectation.params = &RepositoryMockGetListParams{ctx, filter}
	for _, e := range mmGetList.expectations {
		if minimock.Equal(e.params, mmGetList.defaultExpectation.params) {
			mmGetList.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetList.defaultExpectation.params)
		}
	}

	return mmGetList
}

// Inspect accepts an inspector function that has same arguments as the Repository.GetList
func (mmGetList *mRepositoryMockGetList) Inspect(f func(ctx context.Context, filter *mm_student_repository.Query)) *mRepositoryMockGetList {
	if mmGetList.mock.inspectFuncGetList != nil {
		mmGetList.mock.t.Fatalf("Inspect function is already set for RepositoryMock.GetList")
	}

	mmGetList.mock.inspectFuncGetList = f

	return mmGetList
}

// Return sets up results that will be returned by Repository.GetList
func (mmGetList *mRepositoryMockGetList) Return(spa1 []*model.Student, err error) *RepositoryMock {
	if mmGetList.mock.funcGetList != nil {
		mmGetList.mock.t.Fatalf("RepositoryMock.GetList mock is already set by Set")
	}

	if mmGetList.defaultExpectation == nil {
		mmGetList.defaultExpectation = &RepositoryMockGetListExpectation{mock: mmGetList.mock}
	}
	mmGetList.defaultExpectation.results = &RepositoryMockGetListResults{spa1, err}
	return mmGetList.mock
}

//Set uses given function f to mock the Repository.GetList method
func (mmGetList *mRepositoryMockGetList) Set(f func(ctx context.Context, filter *mm_student_repository.Query) (spa1 []*model.Student, err error)) *RepositoryMock {
	if mmGetList.defaultExpectation != nil {
		mmGetList.mock.t.Fatalf("Default expectation is already set for the Repository.GetList method")
	}

	if len(mmGetList.expectations) > 0 {
		mmGetList.mock.t.Fatalf("Some expectations are already set for the Repository.GetList method")
	}

	mmGetList.mock.funcGetList = f
	return mmGetList.mock
}

// When sets expectation for the Repository.GetList which will trigger the result defined by the following
// Then helper
func (mmGetList *mRepositoryMockGetList) When(ctx context.Context, filter *mm_student_repository.Query) *RepositoryMockGetListExpectation {
	if mmGetList.mock.funcGetList != nil {
		mmGetList.mock.t.Fatalf("RepositoryMock.GetList mock is already set by Set")
	}

	expectation := &RepositoryMockGetListExpectation{
		mock:   mmGetList.mock,
		params: &RepositoryMockGetListParams{ctx, filter},
	}
	mmGetList.expectations = append(mmGetList.expectations, expectation)
	return expectation
}

// Then sets up Repository.GetList return parameters for the expectation previously defined by the When method
func (e *RepositoryMockGetListExpectation) Then(spa1 []*model.Student, err error) *RepositoryMock {
	e.results = &RepositoryMockGetListResults{spa1, err}
	return e.mock
}

// GetList implements student_repository.Repository
func (mmGetList *RepositoryMock) GetList(ctx context.Context, filter *mm_student_repository.Query) (spa1 []*model.Student, err error) {
	mm_atomic.AddUint64(&mmGetList.beforeGetListCounter, 1)
	defer mm_atomic.AddUint64(&mmGetList.afterGetListCounter, 1)

	if mmGetList.inspectFuncGetList != nil {
		mmGetList.inspectFuncGetList(ctx, filter)
	}

	mm_params := &RepositoryMockGetListParams{ctx, filter}

	// Record call args
	mmGetList.GetListMock.mutex.Lock()
	mmGetList.GetListMock.callArgs = append(mmGetList.GetListMock.callArgs, mm_params)
	mmGetList.GetListMock.mutex.Unlock()

	for _, e := range mmGetList.GetListMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.spa1, e.results.err
		}
	}

	if mmGetList.GetListMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetList.GetListMock.defaultExpectation.Counter, 1)
		mm_want := mmGetList.GetListMock.defaultExpectation.params
		mm_got := RepositoryMockGetListParams{ctx, filter}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetList.t.Errorf("RepositoryMock.GetList got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetList.GetListMock.defaultExpectation.results
		if mm_results == nil {
			mmGetList.t.Fatal("No results are set for the RepositoryMock.GetList")
		}
		return (*mm_results).spa1, (*mm_results).err
	}
	if mmGetList.funcGetList != nil {
		return mmGetList.funcGetList(ctx, filter)
	}
	mmGetList.t.Fatalf("Unexpected call to RepositoryMock.GetList. %v %v", ctx, filter)
	return
}

// GetListAfterCounter returns a count of finished RepositoryMock.GetList invocations
func (mmGetList *RepositoryMock) GetListAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetList.afterGetListCounter)
}

// GetListBeforeCounter returns a count of RepositoryMock.GetList invocations
func (mmGetList *RepositoryMock) GetListBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetList.beforeGetListCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.GetList.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetList *mRepositoryMockGetList) Calls() []*RepositoryMockGetListParams {
	mmGetList.mutex.RLock()

	argCopy := make([]*RepositoryMockGetListParams, len(mmGetList.callArgs))
	copy(argCopy, mmGetList.callArgs)

	mmGetList.mutex.RUnlock()

	return argCopy
}

// MinimockGetListDone returns true if the count of the GetList invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockGetListDone() bool {
	for _, e := range m.GetListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetListCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetList != nil && mm_atomic.LoadUint64(&m.afterGetListCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetListInspect logs each unmet expectation
func (m *RepositoryMock) MinimockGetListInspect() {
	for _, e := range m.GetListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.GetList with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetListCounter) < 1 {
		if m.GetListMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.GetList")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.GetList with params: %#v", *m.GetListMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetList != nil && mm_atomic.LoadUint64(&m.afterGetListCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.GetList")
	}
}

type mRepositoryMockIsExist struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockIsExistExpectation
	expectations       []*RepositoryMockIsExistExpectation

	callArgs []*RepositoryMockIsExistParams
	mutex    sync.RWMutex
}

// RepositoryMockIsExistExpectation specifies expectation struct of the Repository.IsExist
type RepositoryMockIsExistExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockIsExistParams
	results *RepositoryMockIsExistResults
	Counter uint64
}

// RepositoryMockIsExistParams contains parameters of the Repository.IsExist
type RepositoryMockIsExistParams struct {
	ctx            context.Context
	telegramChatID int64
}

// RepositoryMockIsExistResults contains results of the Repository.IsExist
type RepositoryMockIsExistResults struct {
	b1  bool
	err error
}

// Expect sets up expected params for Repository.IsExist
func (mmIsExist *mRepositoryMockIsExist) Expect(ctx context.Context, telegramChatID int64) *mRepositoryMockIsExist {
	if mmIsExist.mock.funcIsExist != nil {
		mmIsExist.mock.t.Fatalf("RepositoryMock.IsExist mock is already set by Set")
	}

	if mmIsExist.defaultExpectation == nil {
		mmIsExist.defaultExpectation = &RepositoryMockIsExistExpectation{}
	}

	mmIsExist.defaultExpectation.params = &RepositoryMockIsExistParams{ctx, telegramChatID}
	for _, e := range mmIsExist.expectations {
		if minimock.Equal(e.params, mmIsExist.defaultExpectation.params) {
			mmIsExist.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmIsExist.defaultExpectation.params)
		}
	}

	return mmIsExist
}

// Inspect accepts an inspector function that has same arguments as the Repository.IsExist
func (mmIsExist *mRepositoryMockIsExist) Inspect(f func(ctx context.Context, telegramChatID int64)) *mRepositoryMockIsExist {
	if mmIsExist.mock.inspectFuncIsExist != nil {
		mmIsExist.mock.t.Fatalf("Inspect function is already set for RepositoryMock.IsExist")
	}

	mmIsExist.mock.inspectFuncIsExist = f

	return mmIsExist
}

// Return sets up results that will be returned by Repository.IsExist
func (mmIsExist *mRepositoryMockIsExist) Return(b1 bool, err error) *RepositoryMock {
	if mmIsExist.mock.funcIsExist != nil {
		mmIsExist.mock.t.Fatalf("RepositoryMock.IsExist mock is already set by Set")
	}

	if mmIsExist.defaultExpectation == nil {
		mmIsExist.defaultExpectation = &RepositoryMockIsExistExpectation{mock: mmIsExist.mock}
	}
	mmIsExist.defaultExpectation.results = &RepositoryMockIsExistResults{b1, err}
	return mmIsExist.mock
}

//Set uses given function f to mock the Repository.IsExist method
func (mmIsExist *mRepositoryMockIsExist) Set(f func(ctx context.Context, telegramChatID int64) (b1 bool, err error)) *RepositoryMock {
	if mmIsExist.defaultExpectation != nil {
		mmIsExist.mock.t.Fatalf("Default expectation is already set for the Repository.IsExist method")
	}

	if len(mmIsExist.expectations) > 0 {
		mmIsExist.mock.t.Fatalf("Some expectations are already set for the Repository.IsExist method")
	}

	mmIsExist.mock.funcIsExist = f
	return mmIsExist.mock
}

// When sets expectation for the Repository.IsExist which will trigger the result defined by the following
// Then helper
func (mmIsExist *mRepositoryMockIsExist) When(ctx context.Context, telegramChatID int64) *RepositoryMockIsExistExpectation {
	if mmIsExist.mock.funcIsExist != nil {
		mmIsExist.mock.t.Fatalf("RepositoryMock.IsExist mock is already set by Set")
	}

	expectation := &RepositoryMockIsExistExpectation{
		mock:   mmIsExist.mock,
		params: &RepositoryMockIsExistParams{ctx, telegramChatID},
	}
	mmIsExist.expectations = append(mmIsExist.expectations, expectation)
	return expectation
}

// Then sets up Repository.IsExist return parameters for the expectation previously defined by the When method
func (e *RepositoryMockIsExistExpectation) Then(b1 bool, err error) *RepositoryMock {
	e.results = &RepositoryMockIsExistResults{b1, err}
	return e.mock
}

// IsExist implements student_repository.Repository
func (mmIsExist *RepositoryMock) IsExist(ctx context.Context, telegramChatID int64) (b1 bool, err error) {
	mm_atomic.AddUint64(&mmIsExist.beforeIsExistCounter, 1)
	defer mm_atomic.AddUint64(&mmIsExist.afterIsExistCounter, 1)

	if mmIsExist.inspectFuncIsExist != nil {
		mmIsExist.inspectFuncIsExist(ctx, telegramChatID)
	}

	mm_params := &RepositoryMockIsExistParams{ctx, telegramChatID}

	// Record call args
	mmIsExist.IsExistMock.mutex.Lock()
	mmIsExist.IsExistMock.callArgs = append(mmIsExist.IsExistMock.callArgs, mm_params)
	mmIsExist.IsExistMock.mutex.Unlock()

	for _, e := range mmIsExist.IsExistMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmIsExist.IsExistMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmIsExist.IsExistMock.defaultExpectation.Counter, 1)
		mm_want := mmIsExist.IsExistMock.defaultExpectation.params
		mm_got := RepositoryMockIsExistParams{ctx, telegramChatID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmIsExist.t.Errorf("RepositoryMock.IsExist got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmIsExist.IsExistMock.defaultExpectation.results
		if mm_results == nil {
			mmIsExist.t.Fatal("No results are set for the RepositoryMock.IsExist")
		}
		return (*mm_results).b1, (*mm_results).err
	}
	if mmIsExist.funcIsExist != nil {
		return mmIsExist.funcIsExist(ctx, telegramChatID)
	}
	mmIsExist.t.Fatalf("Unexpected call to RepositoryMock.IsExist. %v %v", ctx, telegramChatID)
	return
}

// IsExistAfterCounter returns a count of finished RepositoryMock.IsExist invocations
func (mmIsExist *RepositoryMock) IsExistAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsExist.afterIsExistCounter)
}

// IsExistBeforeCounter returns a count of RepositoryMock.IsExist invocations
func (mmIsExist *RepositoryMock) IsExistBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsExist.beforeIsExistCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.IsExist.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmIsExist *mRepositoryMockIsExist) Calls() []*RepositoryMockIsExistParams {
	mmIsExist.mutex.RLock()

	argCopy := make([]*RepositoryMockIsExistParams, len(mmIsExist.callArgs))
	copy(argCopy, mmIsExist.callArgs)

	mmIsExist.mutex.RUnlock()

	return argCopy
}

// MinimockIsExistDone returns true if the count of the IsExist invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockIsExistDone() bool {
	for _, e := range m.IsExistMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IsExistMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIsExistCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIsExist != nil && mm_atomic.LoadUint64(&m.afterIsExistCounter) < 1 {
		return false
	}
	return true
}

// MinimockIsExistInspect logs each unmet expectation
func (m *RepositoryMock) MinimockIsExistInspect() {
	for _, e := range m.IsExistMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.IsExist with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IsExistMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIsExistCounter) < 1 {
		if m.IsExistMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.IsExist")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.IsExist with params: %#v", *m.IsExistMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIsExist != nil && mm_atomic.LoadUint64(&m.afterIsExistCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.IsExist")
	}
}

type mRepositoryMockUpdate struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockUpdateExpectation
	expectations       []*RepositoryMockUpdateExpectation

	callArgs []*RepositoryMockUpdateParams
	mutex    sync.RWMutex
}

// RepositoryMockUpdateExpectation specifies expectation struct of the Repository.Update
type RepositoryMockUpdateExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockUpdateParams
	results *RepositoryMockUpdateResults
	Counter uint64
}

// RepositoryMockUpdateParams contains parameters of the Repository.Update
type RepositoryMockUpdateParams struct {
	ctx           context.Context
	telegramID    int64
	updateStudent *model.UpdateStudent
}

// RepositoryMockUpdateResults contains results of the Repository.Update
type RepositoryMockUpdateResults struct {
	err error
}

// Expect sets up expected params for Repository.Update
func (mmUpdate *mRepositoryMockUpdate) Expect(ctx context.Context, telegramID int64, updateStudent *model.UpdateStudent) *mRepositoryMockUpdate {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("RepositoryMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &RepositoryMockUpdateExpectation{}
	}

	mmUpdate.defaultExpectation.params = &RepositoryMockUpdateParams{ctx, telegramID, updateStudent}
	for _, e := range mmUpdate.expectations {
		if minimock.Equal(e.params, mmUpdate.defaultExpectation.params) {
			mmUpdate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdate.defaultExpectation.params)
		}
	}

	return mmUpdate
}

// Inspect accepts an inspector function that has same arguments as the Repository.Update
func (mmUpdate *mRepositoryMockUpdate) Inspect(f func(ctx context.Context, telegramID int64, updateStudent *model.UpdateStudent)) *mRepositoryMockUpdate {
	if mmUpdate.mock.inspectFuncUpdate != nil {
		mmUpdate.mock.t.Fatalf("Inspect function is already set for RepositoryMock.Update")
	}

	mmUpdate.mock.inspectFuncUpdate = f

	return mmUpdate
}

// Return sets up results that will be returned by Repository.Update
func (mmUpdate *mRepositoryMockUpdate) Return(err error) *RepositoryMock {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("RepositoryMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &RepositoryMockUpdateExpectation{mock: mmUpdate.mock}
	}
	mmUpdate.defaultExpectation.results = &RepositoryMockUpdateResults{err}
	return mmUpdate.mock
}

//Set uses given function f to mock the Repository.Update method
func (mmUpdate *mRepositoryMockUpdate) Set(f func(ctx context.Context, telegramID int64, updateStudent *model.UpdateStudent) (err error)) *RepositoryMock {
	if mmUpdate.defaultExpectation != nil {
		mmUpdate.mock.t.Fatalf("Default expectation is already set for the Repository.Update method")
	}

	if len(mmUpdate.expectations) > 0 {
		mmUpdate.mock.t.Fatalf("Some expectations are already set for the Repository.Update method")
	}

	mmUpdate.mock.funcUpdate = f
	return mmUpdate.mock
}

// When sets expectation for the Repository.Update which will trigger the result defined by the following
// Then helper
func (mmUpdate *mRepositoryMockUpdate) When(ctx context.Context, telegramID int64, updateStudent *model.UpdateStudent) *RepositoryMockUpdateExpectation {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("RepositoryMock.Update mock is already set by Set")
	}

	expectation := &RepositoryMockUpdateExpectation{
		mock:   mmUpdate.mock,
		params: &RepositoryMockUpdateParams{ctx, telegramID, updateStudent},
	}
	mmUpdate.expectations = append(mmUpdate.expectations, expectation)
	return expectation
}

// Then sets up Repository.Update return parameters for the expectation previously defined by the When method
func (e *RepositoryMockUpdateExpectation) Then(err error) *RepositoryMock {
	e.results = &RepositoryMockUpdateResults{err}
	return e.mock
}

// Update implements student_repository.Repository
func (mmUpdate *RepositoryMock) Update(ctx context.Context, telegramID int64, updateStudent *model.UpdateStudent) (err error) {
	mm_atomic.AddUint64(&mmUpdate.beforeUpdateCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdate.afterUpdateCounter, 1)

	if mmUpdate.inspectFuncUpdate != nil {
		mmUpdate.inspectFuncUpdate(ctx, telegramID, updateStudent)
	}

	mm_params := &RepositoryMockUpdateParams{ctx, telegramID, updateStudent}

	// Record call args
	mmUpdate.UpdateMock.mutex.Lock()
	mmUpdate.UpdateMock.callArgs = append(mmUpdate.UpdateMock.callArgs, mm_params)
	mmUpdate.UpdateMock.mutex.Unlock()

	for _, e := range mmUpdate.UpdateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdate.UpdateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdate.UpdateMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdate.UpdateMock.defaultExpectation.params
		mm_got := RepositoryMockUpdateParams{ctx, telegramID, updateStudent}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdate.t.Errorf("RepositoryMock.Update got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdate.UpdateMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdate.t.Fatal("No results are set for the RepositoryMock.Update")
		}
		return (*mm_results).err
	}
	if mmUpdate.funcUpdate != nil {
		return mmUpdate.funcUpdate(ctx, telegramID, updateStudent)
	}
	mmUpdate.t.Fatalf("Unexpected call to RepositoryMock.Update. %v %v %v", ctx, telegramID, updateStudent)
	return
}

// UpdateAfterCounter returns a count of finished RepositoryMock.Update invocations
func (mmUpdate *RepositoryMock) UpdateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.afterUpdateCounter)
}

// UpdateBeforeCounter returns a count of RepositoryMock.Update invocations
func (mmUpdate *RepositoryMock) UpdateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.beforeUpdateCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.Update.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdate *mRepositoryMockUpdate) Calls() []*RepositoryMockUpdateParams {
	mmUpdate.mutex.RLock()

	argCopy := make([]*RepositoryMockUpdateParams, len(mmUpdate.callArgs))
	copy(argCopy, mmUpdate.callArgs)

	mmUpdate.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateDone returns true if the count of the Update invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockUpdateDone() bool {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateInspect logs each unmet expectation
func (m *RepositoryMock) MinimockUpdateInspect() {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.Update with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		if m.UpdateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.Update")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.Update with params: %#v", *m.UpdateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.Update")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockGetListInspect()

		m.MinimockIsExistInspect()

		m.MinimockUpdateInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *RepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockGetListDone() &&
		m.MinimockIsExistDone() &&
		m.MinimockUpdateDone()
}
