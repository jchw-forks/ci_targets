#ifndef CAST_TO_TOP_H_
#define CAST_TO_TOP_H_

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include <kaitai/kaitaistruct.h>
#include <kaitai/kaitaistream.h>

#include <stdint.h>

#if KAITAI_STRUCT_VERSION < 7000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.7 or later is required"
#endif

class cast_to_top_t : public kaitai::kstruct {

public:

    cast_to_top_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = 0, cast_to_top_t* p__root = 0);
    void _read();
    ~cast_to_top_t();

private:
    bool f_header;
    cast_to_top_t* m_header;

public:
    cast_to_top_t* header();

private:
    bool f_header_casted;
    cast_to_top_t* m_header_casted;

public:
    cast_to_top_t* header_casted();

private:
    uint8_t m_code;
    cast_to_top_t* m__root;
    kaitai::kstruct* m__parent;

public:
    uint8_t code() const { return m_code; }
    cast_to_top_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};

#endif  // CAST_TO_TOP_H_