// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "enum_1.h"

#include <iostream>
#include <fstream>

enum_1_t::enum_1_t(kaitai::kstream *p_io, kaitai::kstruct* p_parent, enum_1_t *p_root) : kaitai::kstruct(p_io) {
    m__parent = p_parent;
    m__root = this;
    m_main = new main_obj_t(m__io, this, m__root);
}

enum_1_t::~enum_1_t() {
    delete m_main;
}

enum_1_t::main_obj_t::main_obj_t(kaitai::kstream *p_io, enum_1_t* p_parent, enum_1_t *p_root) : kaitai::kstruct(p_io) {
    m__parent = p_parent;
    m__root = p_root;
    m_submain = new submain_obj_t(m__io, this, m__root);
}

enum_1_t::main_obj_t::~main_obj_t() {
    delete m_submain;
}

enum_1_t::main_obj_t::submain_obj_t::submain_obj_t(kaitai::kstream *p_io, enum_1_t::main_obj_t* p_parent, enum_1_t *p_root) : kaitai::kstruct(p_io) {
    m__parent = p_parent;
    m__root = p_root;
    m_pet_1 = static_cast<enum_1_t::main_obj_t::animal_t>(m__io->read_u4le());
    m_pet_2 = static_cast<enum_1_t::main_obj_t::animal_t>(m__io->read_u4le());
}

enum_1_t::main_obj_t::submain_obj_t::~submain_obj_t() {
}
