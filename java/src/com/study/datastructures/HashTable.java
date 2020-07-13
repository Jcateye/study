package com.study.datastructures;

public class HashTable {

    private Node[] table;

    private int size;

    private int factor;

    public void put(Object key, Object obj){
        int hash = key.hashCode();
        int index = hash % table.length;

        Node node = new Node(key, obj);

        if(table[index] == null) {
            table[index] = node;
        }
        else if {

            Node head = table[index];



            node.next = head;
        }

        size++;


    }

    class Node {
        private Node next;

        private Object value;


    }
}
