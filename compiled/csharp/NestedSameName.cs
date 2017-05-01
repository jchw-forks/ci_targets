// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

using System;
using System.Collections.Generic;
using System.Linq;

namespace Kaitai
{
    public partial class NestedSameName : KaitaiStruct
    {
        public static NestedSameName FromFile(string fileName)
        {
            return new NestedSameName(new KaitaiStream(fileName));
        }

        public NestedSameName(KaitaiStream io, KaitaiStruct parent = null, NestedSameName root = null) : base(io)
        {
            m_parent = parent;
            m_root = root ?? this;
            _parse();
        }

        private void _parse()
        {
            _mainData = new Main(m_io, this, m_root);
            _dummy = new DummyObj(m_io, this, m_root);
        }
        public partial class Main : KaitaiStruct
        {
            public static Main FromFile(string fileName)
            {
                return new Main(new KaitaiStream(fileName));
            }

            public Main(KaitaiStream io, NestedSameName parent = null, NestedSameName root = null) : base(io)
            {
                m_parent = parent;
                m_root = root;
                _parse();
            }

            private void _parse()
            {
                _mainSize = m_io.ReadS4le();
                _foo = new FooObj(m_io, this, m_root);
            }
            public partial class FooObj : KaitaiStruct
            {
                public static FooObj FromFile(string fileName)
                {
                    return new FooObj(new KaitaiStream(fileName));
                }

                public FooObj(KaitaiStream io, NestedSameName.Main parent = null, NestedSameName root = null) : base(io)
                {
                    m_parent = parent;
                    m_root = root;
                    _parse();
                }

                private void _parse()
                {
                    _data = m_io.ReadBytes((M_Parent.MainSize * 2));
                }
                private byte[] _data;
                private NestedSameName m_root;
                private NestedSameName.Main m_parent;
                public byte[] Data { get { return _data; } }
                public NestedSameName M_Root { get { return m_root; } }
                public NestedSameName.Main M_Parent { get { return m_parent; } }
            }
            private int _mainSize;
            private FooObj _foo;
            private NestedSameName m_root;
            private NestedSameName m_parent;
            public int MainSize { get { return _mainSize; } }
            public FooObj Foo { get { return _foo; } }
            public NestedSameName M_Root { get { return m_root; } }
            public NestedSameName M_Parent { get { return m_parent; } }
        }
        public partial class DummyObj : KaitaiStruct
        {
            public static DummyObj FromFile(string fileName)
            {
                return new DummyObj(new KaitaiStream(fileName));
            }

            public DummyObj(KaitaiStream io, NestedSameName parent = null, NestedSameName root = null) : base(io)
            {
                m_parent = parent;
                m_root = root;
                _parse();
            }

            private void _parse()
            {
            }
            public partial class Foo : KaitaiStruct
            {
                public static Foo FromFile(string fileName)
                {
                    return new Foo(new KaitaiStream(fileName));
                }

                public Foo(KaitaiStream io, KaitaiStruct parent = null, NestedSameName root = null) : base(io)
                {
                    m_parent = parent;
                    m_root = root;
                    _parse();
                }

                private void _parse()
                {
                }
                private NestedSameName m_root;
                private KaitaiStruct m_parent;
                public NestedSameName M_Root { get { return m_root; } }
                public KaitaiStruct M_Parent { get { return m_parent; } }
            }
            private NestedSameName m_root;
            private NestedSameName m_parent;
            public NestedSameName M_Root { get { return m_root; } }
            public NestedSameName M_Parent { get { return m_parent; } }
        }
        private Main _mainData;
        private DummyObj _dummy;
        private NestedSameName m_root;
        private KaitaiStruct m_parent;
        public Main MainData { get { return _mainData; } }
        public DummyObj Dummy { get { return _dummy; } }
        public NestedSameName M_Root { get { return m_root; } }
        public KaitaiStruct M_Parent { get { return m_parent; } }
    }
}
