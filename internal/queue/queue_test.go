package queue

import (
	"math/rand"
	"testing"
)

// func TestQueue_Enqueue(t *testing.T) {
// 	type fields struct {
// 		data  map[int][]byte
// 		front int
// 		back  int
// 	}
// 	type args struct {
// 		data []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    int
// 		wantErr bool
// 	}{
// 		{
// 			name: "should enqueue data to an empty queue",
// 			fields: fields{
// 				data:  map[int][]byte{},
// 				front: -1,
// 			},
// 			args: args{
// 				data: []byte("test"),
// 			},
// 			want:    0,
// 			wantErr: false,
// 		},
// 		{
// 			name: "should enqueue data to an non empty queue",
// 			fields: fields{
// 				data: map[int][]byte{
// 					3: []byte("test3"),
// 					4: []byte("test4"),
// 				},
// 				front: 4,
// 				back:  3,
// 			},
// 			args: args{
// 				data: []byte("test"),
// 			},
// 			want:    5,
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			q := &Queue{
// 				data:  tt.fields.data,
// 				front: tt.fields.front,
// 				back:  tt.fields.back,
// 			}
// 			got, err := q.Enqueue(tt.args.data)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Queue.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("Queue.Enqueue() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestQueue_Dequeue(t *testing.T) {
// 	type fields struct {
// 		data  map[int][]byte
// 		front int
// 		back  int
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    []byte
// 		wantErr bool
// 	}{
// 		{
// 			name: "should return error dequeue on empty queue",
// 			fields: fields{
// 				data:  map[int][]byte{},
// 				front: -1,
// 				back:  0,
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 		{
// 			name: "should return value on dequeue non-empty queue",
// 			fields: fields{
// 				data: map[int][]byte{
// 					1: []byte("test"),
// 				},
// 				front: 1,
// 				back:  1,
// 			},
// 			want:    []byte("test"),
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			q := &Queue{
// 				data:  tt.fields.data,
// 				front: tt.fields.front,
// 				back:  tt.fields.back,
// 			}
// 			got, err := q.Dequeue()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Queue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestQueue_Update(t *testing.T) {
// 	type fields struct {
// 		data  map[int][]byte
// 		front int
// 		back  int
// 	}
// 	type args struct {
// 		id   int
// 		data []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "should throw error if job is not there with given id",
// 			fields: fields{
// 				data: map[int][]byte{
// 					1: []byte("test"),
// 				},
// 				front: 1,
// 				back:  1,
// 			},
// 			args: args{
// 				id:   2,
// 				data: []byte("test2"),
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "should update the data for the given job id",
// 			fields: fields{
// 				data: map[int][]byte{
// 					1: []byte("test"),
// 				},
// 				front: 1,
// 				back:  1,
// 			},
// 			args: args{
// 				id:   1,
// 				data: []byte("test1"),
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			q := &Queue{
// 				data:  tt.fields.data,
// 				front: tt.fields.front,
// 				back:  tt.fields.back,
// 			}
// 			err := q.Update(tt.args.id, tt.args.data)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Queue.Update() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			if err == nil && string(q.data[tt.args.id]) != string(tt.args.data) {
// 				t.Errorf("Data didnt updated")
// 			}
// 		})
// 	}
// }

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func BenchmarkEnqueue128b(b *testing.B) {
	// run the Fib function b.N times
	q := NewQueue()
	data := RandStringBytes(128)
	for n := 0; n < b.N; n++ {
		q.Enqueue(data)
	}
}

func BenchmarkEnqueue1kb(b *testing.B) {
	// run the Fib function b.N times
	q := NewQueue()
	data := RandStringBytes(1024)
	for n := 0; n < b.N; n++ {
		q.Enqueue(data)
	}
}

func BenchmarkEnqueue1MB(b *testing.B) {
	// run the Fib function b.N times
	q := NewQueue()
	data := RandStringBytes(1024 * 1024)
	for n := 0; n < b.N; n++ {
		q.Enqueue(data)
	}
}
