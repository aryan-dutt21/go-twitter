// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"twitter/ent/predicate"
	"twitter/ent/tweet"
	"twitter/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTweet = "Tweet"
	TypeUser  = "User"
)

// TweetMutation represents an operation that mutates the Tweet nodes in the graph.
type TweetMutation struct {
	config
	op            Op
	typ           string
	id            *int
	text          *string
	clearedFields map[string]struct{}
	author        *int
	clearedauthor bool
	done          bool
	oldValue      func(context.Context) (*Tweet, error)
	predicates    []predicate.Tweet
}

var _ ent.Mutation = (*TweetMutation)(nil)

// tweetOption allows management of the mutation configuration using functional options.
type tweetOption func(*TweetMutation)

// newTweetMutation creates new mutation for the Tweet entity.
func newTweetMutation(c config, op Op, opts ...tweetOption) *TweetMutation {
	m := &TweetMutation{
		config:        c,
		op:            op,
		typ:           TypeTweet,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTweetID sets the ID field of the mutation.
func withTweetID(id int) tweetOption {
	return func(m *TweetMutation) {
		var (
			err   error
			once  sync.Once
			value *Tweet
		)
		m.oldValue = func(ctx context.Context) (*Tweet, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Tweet.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTweet sets the old Tweet of the mutation.
func withTweet(node *Tweet) tweetOption {
	return func(m *TweetMutation) {
		m.oldValue = func(context.Context) (*Tweet, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TweetMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TweetMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TweetMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TweetMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Tweet.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetText sets the "text" field.
func (m *TweetMutation) SetText(s string) {
	m.text = &s
}

// Text returns the value of the "text" field in the mutation.
func (m *TweetMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// OldText returns the old "text" field's value of the Tweet entity.
// If the Tweet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TweetMutation) OldText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldText: %w", err)
	}
	return oldValue.Text, nil
}

// ResetText resets all changes to the "text" field.
func (m *TweetMutation) ResetText() {
	m.text = nil
}

// SetAuthorID sets the "author" edge to the User entity by id.
func (m *TweetMutation) SetAuthorID(id int) {
	m.author = &id
}

// ClearAuthor clears the "author" edge to the User entity.
func (m *TweetMutation) ClearAuthor() {
	m.clearedauthor = true
}

// AuthorCleared reports if the "author" edge to the User entity was cleared.
func (m *TweetMutation) AuthorCleared() bool {
	return m.clearedauthor
}

// AuthorID returns the "author" edge ID in the mutation.
func (m *TweetMutation) AuthorID() (id int, exists bool) {
	if m.author != nil {
		return *m.author, true
	}
	return
}

// AuthorIDs returns the "author" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// AuthorID instead. It exists only for internal usage by the builders.
func (m *TweetMutation) AuthorIDs() (ids []int) {
	if id := m.author; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetAuthor resets all changes to the "author" edge.
func (m *TweetMutation) ResetAuthor() {
	m.author = nil
	m.clearedauthor = false
}

// Where appends a list predicates to the TweetMutation builder.
func (m *TweetMutation) Where(ps ...predicate.Tweet) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TweetMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TweetMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Tweet, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TweetMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TweetMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Tweet).
func (m *TweetMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TweetMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.text != nil {
		fields = append(fields, tweet.FieldText)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TweetMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case tweet.FieldText:
		return m.Text()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TweetMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case tweet.FieldText:
		return m.OldText(ctx)
	}
	return nil, fmt.Errorf("unknown Tweet field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TweetMutation) SetField(name string, value ent.Value) error {
	switch name {
	case tweet.FieldText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetText(v)
		return nil
	}
	return fmt.Errorf("unknown Tweet field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TweetMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TweetMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TweetMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Tweet numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TweetMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TweetMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TweetMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Tweet nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TweetMutation) ResetField(name string) error {
	switch name {
	case tweet.FieldText:
		m.ResetText()
		return nil
	}
	return fmt.Errorf("unknown Tweet field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TweetMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.author != nil {
		edges = append(edges, tweet.EdgeAuthor)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TweetMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case tweet.EdgeAuthor:
		if id := m.author; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TweetMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TweetMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TweetMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedauthor {
		edges = append(edges, tweet.EdgeAuthor)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TweetMutation) EdgeCleared(name string) bool {
	switch name {
	case tweet.EdgeAuthor:
		return m.clearedauthor
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TweetMutation) ClearEdge(name string) error {
	switch name {
	case tweet.EdgeAuthor:
		m.ClearAuthor()
		return nil
	}
	return fmt.Errorf("unknown Tweet unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TweetMutation) ResetEdge(name string) error {
	switch name {
	case tweet.EdgeAuthor:
		m.ResetAuthor()
		return nil
	}
	return fmt.Errorf("unknown Tweet edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	username      *string
	clearedFields map[string]struct{}
	tweets        map[int]struct{}
	removedtweets map[int]struct{}
	clearedtweets bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
}

// AddTweetIDs adds the "tweets" edge to the Tweet entity by ids.
func (m *UserMutation) AddTweetIDs(ids ...int) {
	if m.tweets == nil {
		m.tweets = make(map[int]struct{})
	}
	for i := range ids {
		m.tweets[ids[i]] = struct{}{}
	}
}

// ClearTweets clears the "tweets" edge to the Tweet entity.
func (m *UserMutation) ClearTweets() {
	m.clearedtweets = true
}

// TweetsCleared reports if the "tweets" edge to the Tweet entity was cleared.
func (m *UserMutation) TweetsCleared() bool {
	return m.clearedtweets
}

// RemoveTweetIDs removes the "tweets" edge to the Tweet entity by IDs.
func (m *UserMutation) RemoveTweetIDs(ids ...int) {
	if m.removedtweets == nil {
		m.removedtweets = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.tweets, ids[i])
		m.removedtweets[ids[i]] = struct{}{}
	}
}

// RemovedTweets returns the removed IDs of the "tweets" edge to the Tweet entity.
func (m *UserMutation) RemovedTweetsIDs() (ids []int) {
	for id := range m.removedtweets {
		ids = append(ids, id)
	}
	return
}

// TweetsIDs returns the "tweets" edge IDs in the mutation.
func (m *UserMutation) TweetsIDs() (ids []int) {
	for id := range m.tweets {
		ids = append(ids, id)
	}
	return
}

// ResetTweets resets all changes to the "tweets" edge.
func (m *UserMutation) ResetTweets() {
	m.tweets = nil
	m.clearedtweets = false
	m.removedtweets = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUsername:
		return m.Username()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldUsername:
		return m.OldUsername(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.tweets != nil {
		edges = append(edges, user.EdgeTweets)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTweets:
		ids := make([]ent.Value, 0, len(m.tweets))
		for id := range m.tweets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtweets != nil {
		edges = append(edges, user.EdgeTweets)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTweets:
		ids := make([]ent.Value, 0, len(m.removedtweets))
		for id := range m.removedtweets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtweets {
		edges = append(edges, user.EdgeTweets)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeTweets:
		return m.clearedtweets
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeTweets:
		m.ResetTweets()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
